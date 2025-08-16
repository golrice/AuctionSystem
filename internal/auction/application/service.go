package application

import (
	wsDomain "auctionsystem/api/route/ws/domain"
	auctionDomain "auctionsystem/internal/auction/domain"
	"auctionsystem/internal/auction/shared"
	"context"
	"errors"
	"fmt"
	"time"
)

type AuctionService struct {
	auctionRepo    auctionDomain.AuctionRepository
	contextTimeout time.Duration

	publishQ wsDomain.AuctionMessageRepository
}

func NewAuctionService(auctionRepo auctionDomain.AuctionRepository, contextTimeout time.Duration, q wsDomain.AuctionMessageRepository) *AuctionService {
	return &AuctionService{
		auctionRepo:    auctionRepo,
		contextTimeout: contextTimeout,

		publishQ: q,
	}
}

// 创建拍卖
// params:
// cmd: 创建拍卖的命令 包括了拍卖品的标题、描述、开始时间、结束时间、初始价格、步长
// return: 错误信息
func (s *AuctionService) CreateAuction(cmd *CreateAuctionCommand) error {
	// 无法确定某一次拍卖是否之前提供过的 所以这里默认是可以创建的
	// 同名和同样描述并不能说明拍卖物品是一致的 所以是合法的
	if cmd.Title == "" || cmd.Description == "" {
		return errors.New("title or description is empty")
	}
	if cmd.StartTime < time.Now().Unix() {
		return errors.New("start time is in the past")
	}
	auction := &auctionDomain.Auction{
		UserID:      cmd.UserID,
		Title:       cmd.Title,
		Description: cmd.Description,
		StartTime:   cmd.StartTime,
		EndTime:     cmd.EndTime,
		InitPrice:   cmd.InitPrice,
		Step:        cmd.Step,
		Status:      shared.AuctionStatusWaiting,
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()
	return s.auctionRepo.CreateAuction(ctx, auction)
}

// 查看拍卖品
// params:
// query: 查看拍卖品的查询参数 包括了分页信息
// return: 拍卖品的简要信息列表 和 错误信息
func (s *AuctionService) ListAuctions(query *ListAuctionsQuery) ([]*AuctionBriefDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()
	auctions, err := s.auctionRepo.FindAuctions(ctx, query.Pagination)
	if err != nil {
		return nil, err
	}
	return convertToBriefAuctionDTOs(auctions), nil
}

// 获取拍卖品详情
// params:
// query: 获取拍卖品详情的查询参数 包括了拍卖品的ID
// return: 拍卖品的详细信息 和 错误信息
func (s *AuctionService) GetAuctionDetail(query *GetAuctionDetailQuery) (*AuctionDetailDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()
	auction, err := s.auctionRepo.FindAuctionByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	return convertToDetailAuctionDTO(auction), nil
}

// 开始一次拍卖 在时间到了startTime之后异步开启
// params:
// cmd: 开始拍卖的命令 包括了拍卖品的ID
// return: 错误信息
func (s *AuctionService) StartAuction(cmd *StartAuctionCommand) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()
	auction, err := s.auctionRepo.FindAuctionByID(ctx, cmd.AuctionID)
	if err != nil {
		return err
	}
	if auction.Status != shared.AuctionStatusWaiting {
		return errors.New("auction is not waiting")
	}
	auction.Status = shared.AuctionStatusRunning
	return s.auctionRepo.UpdateAuction(ctx, auction)
}

// 结束一次拍卖 在时间到了endTime之后异步结束
// params:
// cmd: 结束拍卖的命令 包括了拍卖品的ID
// return: 错误信息
func (s *AuctionService) EndAuction(cmd *EndAuctionCommand) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()
	auction, err := s.auctionRepo.FindAuctionByID(ctx, cmd.AuctionID)
	if err != nil {
		return err
	}
	if auction.Status != shared.AuctionStatusRunning {
		return errors.New("auction is not running")
	}
	auction.Status = shared.AuctionStatusEnded
	return s.auctionRepo.UpdateAuction(ctx, auction)
}

// 创建出价
// params:
// cmd: 创建出价的命令 包括了拍卖品的ID、用户ID、出价金额
// return: 错误信息
func (s *AuctionService) CreateBid(cmd *CreateBidCommand) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()

	lockKey := fmt.Sprintf("auction:%d:lock", cmd.AuctionID)
	if err := s.auctionRepo.Lock(ctx, lockKey); err != nil {
		return err
	}

	auction, err := s.auctionRepo.FindAuctionByID(ctx, cmd.AuctionID)
	if err != nil {
		return err
	}
	if auction.Status != shared.AuctionStatusRunning {
		return errors.New("auction is not running")
	}
	highestBid, err := s.auctionRepo.LoadAuctionLatestBid(ctx, auction)
	if err != nil {
		return errors.New("load auction latest bids failed")
	}
	bid, err := auction.CreateValidBid(cmd.UserID, cmd.Price, highestBid)
	if err != nil {
		return err
	}
	if err := s.auctionRepo.CreateBid(ctx, bid); err != nil {
		return err
	}

	if err := s.auctionRepo.Unlock(ctx, lockKey); err != nil {
		return err
	}

	return s.publishQ.Publish(ctx, wsDomain.AuctionMessage{
		AuctionID: cmd.AuctionID,
		BidInfo: wsDomain.BidInfo{
			Price:     int(cmd.Price),
			Timestamp: uint(time.Now().Unix()),
		},
	})
}

// 查看最近出价
// params:
// query: 查看最近出价的查询参数 包括了拍卖品的ID、分页信息
// return: 出价的列表 和 错误信息
func (s *AuctionService) ListLatestBids(query *ListLatestBidsQuery) ([]*BidBriefDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()
	auction, err := s.auctionRepo.FindAuctionByID(ctx, query.AuctionID)
	if err != nil {
		return nil, err
	}
	if query.Page == 0 {
		query.Page = 1
	}
	if query.Size == 0 {
		query.Size = 10
	}
	auctionBids, err := s.auctionRepo.LoadAuctionLatestBids(ctx, auction, query.Pagination)
	if err != nil {
		return nil, errors.New("load auction latest bids failed")
	}
	bids := convertToBriefBidDTOs(auctionBids)
	return bids, nil
}

// 查看最高出价
// params:
// query: 查看最高出价的查询参数 包括了拍卖品的ID
// return: 最高出价的信息 和 错误信息
func (s *AuctionService) GetHighestBid(query *GetHighestBidQuery) (*BidBriefDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()
	auction, err := s.auctionRepo.FindAuctionByID(ctx, query.AuctionID)
	if err != nil {
		return nil, err
	}
	highestBid, err := s.auctionRepo.LoadAuctionLatestBid(ctx, auction)
	if err != nil {
		return nil, errors.New("load auction latest bids failed")
	}
	return convertToBriefBidDTO(highestBid), nil
}

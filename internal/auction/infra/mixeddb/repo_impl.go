package mixeddb

import (
	"auctionsystem/internal/auction/domain"
	"auctionsystem/pkg/kernal"
	"context"
)

type AuctionFullRepository struct {
	cache   domain.AuctionCache
	storage domain.AuctionRepository
}

func NewAuctionFullRepository(cache domain.AuctionCache, storage domain.AuctionRepository) domain.AuctionRepository {
	return &AuctionFullRepository{
		cache:   cache,
		storage: storage,
	}
}

func (a *AuctionFullRepository) CreateAuction(ctx context.Context, auction *domain.Auction) error {
	// 直接写入数据库
	return a.storage.CreateAuction(ctx, auction)
}

func (a *AuctionFullRepository) FindAuctionByID(ctx context.Context, id uint) (*domain.Auction, error) {
	// 首先查看缓存
	auction, err := a.cache.FindAuctionByID(ctx, id)
	if err == nil {
		return auction, nil
	}
	// 缓存中没有 则从数据库中查询
	auction, err = a.storage.FindAuctionByID(ctx, id)
	if err != nil {
		return nil, err
	}
	// 写入缓存中
	if err := a.cache.CreateAuction(ctx, auction); err != nil {
		return nil, err
	}
	return auction, nil
}

func (a *AuctionFullRepository) FindAuctions(ctx context.Context, page kernal.Pagination) ([]*domain.Auction, error) {
	// 从数据库中查询
	return a.storage.FindAuctions(ctx, page)
}

func (a *AuctionFullRepository) UpdateAuction(ctx context.Context, auction *domain.Auction) error {
	// 先更新数据库
	if err := a.storage.UpdateAuction(ctx, auction); err != nil {
		return err
	}
	// 再删除缓存
	return a.cache.DeleteAuction(ctx, auction.ID)
}

func (a *AuctionFullRepository) DeleteAuction(ctx context.Context, id uint) error {
	// 先删除数据库
	if err := a.storage.DeleteAuction(ctx, id); err != nil {
		return err
	}
	// 再删除缓存
	return a.cache.DeleteAuction(ctx, id)
}

func (a *AuctionFullRepository) LoadAuctionLatestBids(ctx context.Context, auction *domain.Auction, page kernal.Pagination) ([]*domain.Bid, error) {
	// 从数据库中查询
	return a.storage.LoadAuctionLatestBids(ctx, auction, page)
}

func (a *AuctionFullRepository) LoadAuctionLatestBid(ctx context.Context, auction *domain.Auction) (*domain.Bid, error) {
	// 先从缓存中查询
	bid, err := a.cache.LoadAuctionLatestBid(ctx, auction)
	if err == nil {
		return bid, nil
	}

	// 缓存中没有 则从数据库中查询
	bid, err = a.storage.LoadAuctionLatestBid(ctx, auction)
	if err != nil {
		return nil, err
	}
	// 写入缓存中
	if err := a.cache.CreateBid(ctx, bid); err != nil {
		return nil, err
	}
	return bid, nil
}

func (a *AuctionFullRepository) CreateBid(ctx context.Context, bid *domain.Bid) error {
	// 直接写入数据库
	if err := a.storage.CreateBid(ctx, bid); err != nil {
		return err
	}
	// 写入缓存中
	if err := a.cache.CreateBid(ctx, bid); err != nil {
		return err
	}
	return nil
}

var _ domain.AuctionRepository = (*AuctionFullRepository)(nil)

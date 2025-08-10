package persistence

import (
	"auctionsystem/internal/auction/domain"
	"auctionsystem/internal/auction/infra/adaptor"
	"auctionsystem/pkg/kernal"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type AuctionRepositoryImpl struct {
	db *gorm.DB
}

func NewAuctionRepositoryImpl(db *gorm.DB) domain.AuctionRepository {
	return &AuctionRepositoryImpl{db: db}
}

func (a *AuctionRepositoryImpl) CreateAuction(ctx context.Context, auction *domain.Auction) error {
	if err := a.db.Create(adaptor.ConvertToAuctionModel(auction)).Error; err != nil {
		return err
	}
	return nil
}

func (a *AuctionRepositoryImpl) FindAuctionByID(ctx context.Context, id uint) (*domain.Auction, error) {
	var auctionModel adaptor.AuctionModel
	if err := a.db.Where("id = ?", id).First(&auctionModel).Error; err != nil {
		return nil, err
	}
	return adaptor.ConvertToDomainAuction(&auctionModel), nil
}

func (a *AuctionRepositoryImpl) FindAuctions(ctx context.Context, page kernal.Pagination) ([]*domain.Auction, error) {
	var auctionModels []*adaptor.AuctionModel
	fmt.Println(page)
	if err := a.db.
		Limit(int(page.Limit())).
		Offset(int(page.Offset())).
		Find(&auctionModels).Error; err != nil {
		return nil, err
	}
	return adaptor.ConvertToDomainAuctions(auctionModels), nil
}

func (a *AuctionRepositoryImpl) UpdateAuction(ctx context.Context, auction *domain.Auction) error {
	if err := a.db.Save(adaptor.ConvertToAuctionModel(auction)).Error; err != nil {
		return err
	}
	return nil
}

func (a *AuctionRepositoryImpl) DeleteAuction(ctx context.Context, id uint) error {
	if err := a.db.Delete(&adaptor.AuctionModel{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (a *AuctionRepositoryImpl) LoadAuctionLatestBids(ctx context.Context, auction *domain.Auction, page kernal.Pagination) ([]*domain.Bid, error) {
	var bidModels []adaptor.BidModel
	if err := a.db.Where("auction_id = ?", auction.ID).
		Order("price desc").
		Limit(int(page.Limit())).
		Offset(int(page.Offset())).
		Find(&bidModels).Error; err != nil {
		return nil, err
	}
	return adaptor.ConvertToDomainBids(&bidModels), nil
}

func (a *AuctionRepositoryImpl) LoadAuctionLatestBid(ctx context.Context, auction *domain.Auction) (*domain.Bid, error) {
	var bidModels adaptor.BidModel
	if err := a.db.Where("auction_id = ?", auction.ID).
		Order("price desc").
		First(&bidModels).Error; err != nil {
		return nil, err
	}
	return adaptor.ConvertToDomainBid(&bidModels), nil
}

func (a *AuctionRepositoryImpl) CreateBid(ctx context.Context, bid *domain.Bid) error {
	if err := a.db.Create(adaptor.ConvertToBidModel(bid)).Error; err != nil {
		return err
	}
	return nil
}

var _ domain.AuctionRepository = (*AuctionRepositoryImpl)(nil)

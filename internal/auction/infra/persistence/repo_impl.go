package persistence

import (
	"auctionsystem/internal/auction/domain"
	"auctionsystem/internal/auction/infra/adaptor"
	"auctionsystem/pkg/kernal"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type AuctionPersistencyImpl struct {
	db *gorm.DB
}

func NewAuctionPersistencyImpl(db *gorm.DB) domain.AuctionPersistency {
	return &AuctionPersistencyImpl{db: db}
}

func (a *AuctionPersistencyImpl) CreateAuction(ctx context.Context, auction *domain.Auction) error {
	if err := a.db.Create(adaptor.ConvertToAuctionModel(auction)).Error; err != nil {
		return err
	}
	return nil
}

func (a *AuctionPersistencyImpl) FindAuctionByID(ctx context.Context, id uint) (*domain.Auction, error) {
	var auctionModel adaptor.AuctionModel
	if err := a.db.Where("id = ?", id).First(&auctionModel).Error; err != nil {
		return nil, err
	}
	return adaptor.ConvertToDomainAuction(&auctionModel), nil
}

func (a *AuctionPersistencyImpl) FindAuctions(ctx context.Context, page kernal.Pagination) ([]*domain.Auction, error) {
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

func (a *AuctionPersistencyImpl) UpdateAuction(ctx context.Context, auction *domain.Auction) error {
	if err := a.db.Save(adaptor.ConvertToAuctionModel(auction)).Error; err != nil {
		return err
	}
	return nil
}

func (a *AuctionPersistencyImpl) DeleteAuction(ctx context.Context, id uint) error {
	if err := a.db.Delete(&adaptor.AuctionModel{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (a *AuctionPersistencyImpl) LoadAuctionLatestBids(ctx context.Context, auction *domain.Auction, page kernal.Pagination) ([]*domain.Bid, error) {
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

func (a *AuctionPersistencyImpl) LoadAuctionLatestBid(ctx context.Context, auction *domain.Auction) (*domain.Bid, error) {
	var bidModels adaptor.BidModel
	if err := a.db.Where("auction_id = ?", auction.ID).
		Order("price desc").
		First(&bidModels).Error; err != nil {
		return nil, err
	}
	return adaptor.ConvertToDomainBid(&bidModels), nil
}

func (a *AuctionPersistencyImpl) CreateBid(ctx context.Context, bid *domain.Bid) error {
	if err := a.db.Create(adaptor.ConvertToBidModel(bid)).Error; err != nil {
		return err
	}
	return nil
}

var _ domain.AuctionPersistency = (*AuctionPersistencyImpl)(nil)

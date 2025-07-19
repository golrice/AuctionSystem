package auction

import "gorm.io/gorm"

type auctionRepository struct {
}

type AuctionRepository interface {
	Create(db *gorm.DB, model *AuctionModel) error
}

func NewAuctionRepository() AuctionRepository {
	return &auctionRepository{}
}

func (r *auctionRepository) Create(db *gorm.DB, model *AuctionModel) error {
	return db.Create(model).Error
}

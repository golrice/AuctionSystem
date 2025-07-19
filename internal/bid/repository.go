package bid

import "gorm.io/gorm"

type bidRepository struct {
}

type BidRepository interface {
	Create(db *gorm.DB, model *BidModel) error
}

func NewBidRepository() BidRepository {
	return &bidRepository{}
}

func (r *bidRepository) Create(db *gorm.DB, model *BidModel) error {
	return db.Create(model).Error
}

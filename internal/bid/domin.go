package bid

import "gorm.io/gorm"

type BidRepository interface {
	Create(db *gorm.DB, model *BidModel) error
}

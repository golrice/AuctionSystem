package bid

import "gorm.io/gorm"

type BidStatus int

const (
	BidStatusRunning BidStatus = 1
	BidStatusEnded   BidStatus = 2
)

type BidModel struct {
	gorm.Model
	AuctionID int64     `json:"auction_id" gorm:"column:auction_id"`
	UserID    int64     `json:"user_id" gorm:"column:user_id"`
	Price     int64     `json:"price" gorm:"column:price"`
	Status    BidStatus `json:"status" gorm:"column:status"`
}

func (BidModel) TableName() string {
	return "bids"
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&BidModel{})
}

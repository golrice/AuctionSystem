package bid

import "gorm.io/gorm"

type Bid struct {
	gorm.Model
	AuctionID int64 `json:"auction_id" gorm:"column:auction_id"`
	UserID    int64 `json:"user_id" gorm:"column:user_id"`
	Price     int64 `json:"price" gorm:"column:price"`
	Status    int   `json:"status" gorm:"column:status"`
}

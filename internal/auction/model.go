package auction

import "gorm.io/gorm"

type AuctionStatus int

const (
	AuctionStatusRunning AuctionStatus = 1
	AuctionStatusEnded   AuctionStatus = 2
)

type Auction struct {
	gorm.Model
	Title     string        `json:"title" gorm:"column:title"`
	StartTime int64         `json:"start_time" gorm:"column:start_time"`
	EndTime   int64         `json:"end_time" gorm:"column:end_time"`
	InitPrice int64         `json:"init_price" gorm:"column:init_price"`
	Step      int64         `json:"step" gorm:"column:step"`
	Status    AuctionStatus `json:"status" gorm:"column:status"`
}

package persistence

import (
	"auctionsystem/internal/auction/shared"

	"gorm.io/gorm"
)

type AuctionModel struct {
	gorm.Model
	Title       string               `json:"title" gorm:"column:title; default:''; not null"`
	Description string               `json:"description" gorm:"column:description; default:''; not null"`
	StartTime   int64                `json:"start_time" gorm:"column:start_time; default:0; not null"`
	EndTime     int64                `json:"end_time" gorm:"column:end_time; default:0; not null"`
	InitPrice   int64                `json:"init_price" gorm:"column:init_price; default:0; not null"`
	Step        int64                `json:"step" gorm:"column:step; default:0; not null"`
	Status      shared.AuctionStatus `json:"status" gorm:"column:status; default:0; not null"`

	UserID uint `json:"user_id" gorm:"column:user_id; not null"`
}

func (AuctionModel) TableName() string {
	return "auctions"
}

type BidModel struct {
	gorm.Model

	AuctionID uint  `json:"auction_id" gorm:"column:auction_id; not null"`
	UserID    uint  `json:"user_id" gorm:"column:user_id; not null"`
	Price     int64 `json:"price" gorm:"column:price; default:0; not null"`
}

func (BidModel) TableName() string {
	return "bids"
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&AuctionModel{})
	db.AutoMigrate(&BidModel{})
}

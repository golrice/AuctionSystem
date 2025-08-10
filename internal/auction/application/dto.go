package application

import (
	"auctionsystem/internal/auction/shared"
)

type AuctionBriefDTO struct {
	ID        uint                 `json:"id"`
	Title     string               `json:"title"`
	StartTime int64                `json:"start_time"`
	EndTime   int64                `json:"end_time"`
	InitPrice int64                `json:"init_price"`
	Status    shared.AuctionStatus `json:"status"`
}

type AuctionDetailDTO struct {
	AuctionBriefDTO
	Description string `json:"description"`
	Step        int64  `json:"step"`
}

type BidBriefDTO struct {
	ID    uint  `json:"id"`
	Price int64 `json:"price"`
}

type BidDetailDTO struct {
	BidBriefDTO
	AuctionID uint `json:"auction_id"`
	UserID    uint `json:"user_id"`
}

type AuctionBriefListDTO []*AuctionBriefDTO

type BidBriefListDTO []*BidBriefDTO

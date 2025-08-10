package domain

type AuctionMessage struct {
	AuctionID uint    `json:"auction_id"`
	BidInfo   BidInfo `json:"bid_info"`
}

type BidInfo struct {
	Price     int  `json:"price"`
	Timestamp uint `json:"timestamp"`
}

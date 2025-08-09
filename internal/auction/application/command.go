package application

type CreateAuctionCommand struct {
	UserID      uint
	Title       string
	Description string
	StartTime   int64
	EndTime     int64
	InitPrice   int64
	Step        int64
}

type StartAuctionCommand struct {
	AuctionID uint
	UserID    uint
}

type EndAuctionCommand struct {
	AuctionID uint
	UserID    uint
}

type CreateBidCommand struct {
	AuctionID uint
	UserID    uint
	Price     int64
}

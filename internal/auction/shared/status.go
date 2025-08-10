package shared

type AuctionStatus int

const (
	AuctionStatusRunning AuctionStatus = 1
	AuctionStatusEnded   AuctionStatus = 2
	AuctionStatusWaiting AuctionStatus = 3
)

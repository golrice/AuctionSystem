package application

import "net/http"

type UpgradeCommand struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

type SendBidCommand struct {
	AuctionID uint
	Price     int
	Client    *Client
}

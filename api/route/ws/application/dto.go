package application

import (
	"auctionsystem/api/route/ws/domain"

	"github.com/gorilla/websocket"
)

type WSAuctionMessage struct {
	AuctionID uint    `json:"auction_id"`
	BidInfo   BidInfo `json:"bid_info"`
}

type BidInfo struct {
	Price int `json:"price"`
}

type Client struct {
	AuctionID uint
	Conn      *websocket.Conn
	send      chan<- domain.AuctionMessage
}

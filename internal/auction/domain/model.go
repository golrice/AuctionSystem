package domain

import (
	"auctionsystem/internal/auction/shared"
	"errors"
)

type Auction struct {
	ID uint

	UserID      uint
	Title       string
	Description string
	StartTime   int64
	EndTime     int64
	InitPrice   int64
	Step        int64
	Status      shared.AuctionStatus

	Bids []*Bid
}

type Bid struct {
	ID uint

	AuctionID uint
	UserID    uint
	Price     int64
}

func (a *Auction) CreateBid(userID uint, price int64) (*Bid, error) {
	// 合法化检查
	if len(a.Bids) == 0 && price <= a.InitPrice+a.Step {
		return nil, errors.New("price is less than (current price + smallest step)")
	}
	if price <= a.Bids[len(a.Bids)-1].Price+a.Step {
		return nil, errors.New("price is less than (current price + smallest step)")
	}

	bid := &Bid{
		AuctionID: a.ID,
		UserID:    userID,
		Price:     price,
	}
	a.Bids = append(a.Bids, bid)
	return bid, nil
}

func (a *Auction) IsStarting() bool {
	return a.Status == shared.AuctionStatusRunning
}

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
}

type Bid struct {
	ID uint

	AuctionID uint
	UserID    uint
	Price     int64
}

func (a *Auction) IsStarting() bool {
	return a.Status == shared.AuctionStatusRunning
}

func (a *Auction) CreateValidBid(userID uint, price int64, highestBid *Bid) (*Bid, error) {
	if highestBid == nil {
		if price <= a.InitPrice+a.Step {
			return nil, errors.New("price is less than init price + step")
		}
		return &Bid{
			AuctionID: a.ID,
			UserID:    userID,
			Price:     price,
		}, nil
	}
	if price <= highestBid.Price+a.Step {
		return nil, errors.New("price is less than highest bid price + step")
	}
	return &Bid{
		AuctionID: a.ID,
		UserID:    userID,
		Price:     price,
	}, nil
}

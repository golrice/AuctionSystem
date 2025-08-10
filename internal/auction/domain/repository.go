package domain

import (
	"auctionsystem/pkg/kernal"
	"context"
)

type AuctionPersistency interface {
	CreateAuction(ctx context.Context, auction *Auction) error
	FindAuctionByID(ctx context.Context, id uint) (*Auction, error)
	FindAuctions(ctx context.Context, page kernal.Pagination) ([]*Auction, error)
	UpdateAuction(ctx context.Context, auction *Auction) error
	DeleteAuction(ctx context.Context, id uint) error
	LoadAuctionLatestBids(ctx context.Context, auction *Auction, page kernal.Pagination) ([]*Bid, error)
	LoadAuctionLatestBid(ctx context.Context, auction *Auction) (*Bid, error)

	CreateBid(ctx context.Context, bid *Bid) error
}

type AuctionCache interface {
	CreateAuction(ctx context.Context, auction *Auction) error
	FindAuctionByID(ctx context.Context, id uint) (*Auction, error)
	UpdateAuction(ctx context.Context, auction *Auction) error
	DeleteAuction(ctx context.Context, id uint) error
	LoadAuctionLatestBid(ctx context.Context, auction *Auction) (*Bid, error)

	CreateBid(ctx context.Context, bid *Bid) error
	DeleteBid(ctx context.Context, auctionID uint) error

	Lock(ctx context.Context, key string) error
	Unlock(ctx context.Context, key string) error
}

type AuctionRepository interface {
	CreateAuction(ctx context.Context, auction *Auction) error
	FindAuctionByID(ctx context.Context, id uint) (*Auction, error)
	FindAuctions(ctx context.Context, page kernal.Pagination) ([]*Auction, error)
	UpdateAuction(ctx context.Context, auction *Auction) error
	DeleteAuction(ctx context.Context, id uint) error
	LoadAuctionLatestBids(ctx context.Context, auction *Auction, page kernal.Pagination) ([]*Bid, error)
	LoadAuctionLatestBid(ctx context.Context, auction *Auction) (*Bid, error)

	CreateBid(ctx context.Context, bid *Bid) error

	Lock(ctx context.Context, key string) error
	Unlock(ctx context.Context, key string) error
}

type BidRepository interface {
	FindBidByID(id uint) (*Bid, error)
	FindBidByAuctionID(auctionID uint, page kernal.Pagination) ([]*Bid, error)
	FindBids(page kernal.Pagination) ([]*Bid, error)
	UpdateBid(bid *Bid) error
	DeleteBid(id uint) error
}

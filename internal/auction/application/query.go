package application

import "auctionsystem/pkg/kernal"

type ListAuctionsQuery struct {
	kernal.Pagination
}

type GetAuctionDetailQuery struct {
	ID uint
}

type ListLatestBidsQuery struct {
	kernal.Pagination
	AuctionID uint
}

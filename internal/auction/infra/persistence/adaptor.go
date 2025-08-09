package persistence

import (
	"auctionsystem/internal/auction/domain"
)

func convertToDomainAuction(auction *AuctionModel) *domain.Auction {
	return &domain.Auction{
		ID: auction.ID,

		UserID:      auction.UserID,
		Title:       auction.Title,
		Description: auction.Description,
		StartTime:   auction.StartTime,
		EndTime:     auction.EndTime,
		InitPrice:   auction.InitPrice,
		Step:        auction.Step,
		Status:      auction.Status,

		Bids: []*domain.Bid{},
	}
}

func convertToAuctionModel(auction *domain.Auction) *AuctionModel {
	return &AuctionModel{
		UserID:      auction.UserID,
		Title:       auction.Title,
		Description: auction.Description,
		StartTime:   auction.StartTime,
		EndTime:     auction.EndTime,
		InitPrice:   auction.InitPrice,
		Step:        auction.Step,
		Status:      auction.Status,
	}
}

func convertToDomainAuctions(auctions []*AuctionModel) []*domain.Auction {
	var domainAuctions []*domain.Auction
	for _, auction := range auctions {
		domainAuctions = append(domainAuctions, convertToDomainAuction(auction))
	}
	return domainAuctions
}

func convertToDomainBid(bid *BidModel) *domain.Bid {
	return &domain.Bid{
		ID: bid.ID,

		UserID:    bid.UserID,
		AuctionID: bid.AuctionID,
		Price:     bid.Price,
	}
}

func convertToBidModel(bid *domain.Bid) *BidModel {
	return &BidModel{
		UserID:    bid.UserID,
		AuctionID: bid.AuctionID,
		Price:     bid.Price,
	}
}

func convertToDomainBids(bids *[]BidModel) []*domain.Bid {
	var domainBids []*domain.Bid
	for _, bid := range *bids {
		domainBids = append(domainBids, convertToDomainBid(&bid))
	}
	return domainBids
}

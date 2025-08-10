package adaptor

import (
	"auctionsystem/internal/auction/domain"
)

func ConvertToDomainAuction(auction *AuctionModel) *domain.Auction {
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
	}
}

func ConvertToAuctionModel(auction *domain.Auction) *AuctionModel {
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

func ConvertToDomainAuctions(auctions []*AuctionModel) []*domain.Auction {
	var domainAuctions []*domain.Auction
	for _, auction := range auctions {
		domainAuctions = append(domainAuctions, ConvertToDomainAuction(auction))
	}
	return domainAuctions
}

func ConvertToDomainBid(bid *BidModel) *domain.Bid {
	return &domain.Bid{
		ID: bid.ID,

		UserID:    bid.UserID,
		AuctionID: bid.AuctionID,
		Price:     bid.Price,
	}
}

func ConvertToBidModel(bid *domain.Bid) *BidModel {
	return &BidModel{
		UserID:    bid.UserID,
		AuctionID: bid.AuctionID,
		Price:     bid.Price,
	}
}

func ConvertToDomainBids(bids *[]BidModel) []*domain.Bid {
	var domainBids []*domain.Bid
	for _, bid := range *bids {
		domainBids = append(domainBids, ConvertToDomainBid(&bid))
	}
	return domainBids
}

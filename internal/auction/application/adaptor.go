package application

import "auctionsystem/internal/auction/domain"

func convertToBriefAuctionDTO(auction *domain.Auction) *AuctionBriefDTO {
	return &AuctionBriefDTO{
		ID:        auction.ID,
		Title:     auction.Title,
		StartTime: auction.StartTime,
		EndTime:   auction.EndTime,
		InitPrice: auction.InitPrice,
		Status:    auction.Status,
	}
}

func convertToBriefAuctionDTOs(auctions []*domain.Auction) []*AuctionBriefDTO {
	dtos := make([]*AuctionBriefDTO, 0, len(auctions))
	for _, auction := range auctions {
		dtos = append(dtos, convertToBriefAuctionDTO(auction))
	}
	return dtos
}

func convertToDetailAuctionDTO(auction *domain.Auction) *AuctionDetailDTO {
	return &AuctionDetailDTO{
		AuctionBriefDTO: *convertToBriefAuctionDTO(auction),
		Description:     auction.Description,
		Step:            auction.Step,
	}
}

func convertToBriefBidDTO(bid *domain.Bid) *BidBriefDTO {
	return &BidBriefDTO{
		ID:    bid.ID,
		Price: bid.Price,
	}
}

func convertToBriefBidDTOs(bids []*domain.Bid) []*BidBriefDTO {
	dtos := make([]*BidBriefDTO, 0, len(bids))
	for _, bid := range bids {
		dtos = append(dtos, convertToBriefBidDTO(bid))
	}
	return dtos
}

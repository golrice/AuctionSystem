package bid

import "auctionsystem/internal/common"

type DetailRequestSchema struct {
	AuctionID int64 `json:"auction_id"`
}

type DetailResponseSchema struct {
	common.ResponseSchema
	Bid Description `json:"data"`
}

type CreateRequestSchema struct {
	AuctionID int64 `json:"auction_id"`
	UserID    int64 `json:"user_id"`
	Price     int64 `json:"price"`
}

type CreateResponseSchema struct {
	common.ResponseSchema
}

type UpdateRequestSchema struct {
	AuctionID int64 `json:"auction_id"`
	Price     int64 `json:"price"`
}

type UpdateResponseSchema struct {
	common.ResponseSchema
}

type DeleteRequestSchema struct {
	AuctionID int64 `json:"auction_id"`
}

type DeleteResponseSchema struct {
	common.ResponseSchema
}

type ListRequestSchema struct {
	AuctionID int64 `json:"auction_id"`
}

type ListResponseSchema struct {
	common.ResponseSchema
	Bids []Description `json:"data"`
}

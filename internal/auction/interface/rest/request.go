package rest

import "auctionsystem/pkg/kernal"

// 暴露给用户的接口包括了
// 1. 创建/删除/查看/修改某个拍卖品
// 2. 查看所有拍卖品
// 3. 创建/查看某个拍卖品对应的出价
// 4. 查看某个拍卖品的所有出价

type CreateAuctionRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	InitPrice   int64  `json:"init_price"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	StepPrice   int64  `json:"step_price"`
}

type DeleteAuctionRequest struct {
	ID uint `json:"id"`
}

type GetAuctionRequest struct {
	ID uint `form:"id" json:"id"`
}

type UpdateAuctionRequest struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	StepPrice int64  `json:"step_price"`
}

type ListAuctionRequest struct {
	kernal.Pagination
}

type CreateBidRequest struct {
	AuctionID uint  `json:"auction_id"`
	Price     int64 `json:"price"`
}

type GetHigestBidRequest struct {
	AuctionID uint `form:"auction_id" json:"auction_id" validate:"required"`
}

type ListBidRequest struct {
	AuctionID uint `form:"auction_id" json:"auction_id"`
	kernal.Pagination
}

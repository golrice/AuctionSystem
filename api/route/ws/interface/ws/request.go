package ws

type WSUpgradeRequest struct {
	AuctionID uint `form:"auction_id" binding:"required"`
}

package ws

import (
	wsApplication "auctionsystem/api/route/ws/application"
	auctionApplication "auctionsystem/internal/auction/application"
	"auctionsystem/pkg/kernal"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuctionWSHandler struct {
	hubService     *wsApplication.AuctionHubService
	auctionService *auctionApplication.AuctionService
}

func NewAuctionWSHandler(hubService *wsApplication.AuctionHubService, auctionService *auctionApplication.AuctionService) *AuctionWSHandler {
	return &AuctionWSHandler{
		hubService:     hubService,
		auctionService: auctionService,
	}
}

// ServeWS 处理ws连接
// @Summary 处理ws连接
// @Description 处理ws连接
// @Tags ws
// @Accept json
// @Produce json
// @Param auction_id query int true "拍卖id"
// @Failure 400 {object} kernal.ErrorResult
// @Failure 500 {object} kernal.ErrorResult
// @Router /ws [get]
func (h *AuctionWSHandler) ServeWS(c *gin.Context) {
	var req WSUpgradeRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, kernal.ErrorResult{
			Code: http.StatusBadRequest,
			Msg:  "auction_id is required",
		})
		return
	}

	conn, err := h.hubService.Upgrade(wsApplication.UpgradeCommand{
		Writer:  c.Writer,
		Request: c.Request,
	})
	if err != nil {
		c.JSON(http.StatusOK, kernal.ErrorResult{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	}

	client := h.hubService.RegisterClient(uint(req.AuctionID), conn)

	bids, err := h.auctionService.ListLatestBids(&auctionApplication.ListLatestBidsQuery{
		AuctionID: uint(req.AuctionID),
	})
	if err != nil {
		c.JSON(http.StatusOK, kernal.ErrorResult{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	}

	price := 0
	if len(bids) > 0 {
		price = int(bids[0].Price)
	}

	h.hubService.SendFirstBid(wsApplication.SendBidCommand{
		AuctionID: uint(req.AuctionID),
		Price:     price,
		Client:    client,
	})
}

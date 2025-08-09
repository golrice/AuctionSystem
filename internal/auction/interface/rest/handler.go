package rest

// 暴露给用户的接口包括了
// 1. 创建/删除/查看/修改某个拍卖品
// 2. 查看所有拍卖品
// 3. 创建/查看某个拍卖品对应的出价
// 4. 查看某个拍卖品的所有出价

import (
	"auctionsystem/internal/auction/application"
	"auctionsystem/pkg/kernal"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuctionHandler struct {
	auctionService application.AuctionService
}

func NewAuctionHandler(auctionService application.AuctionService) *AuctionHandler {
	return &AuctionHandler{
		auctionService: auctionService,
	}
}

// @Summary 创建拍卖品
// @Description 创建拍卖品
// @Tags 拍卖品
// @Accept json
// @Produce json
// @Param req body CreateAuctionRequest true "创建拍卖品请求参数"
// @Success 200 {object} kernal.SuccessResult "创建成功"
// @Failure 400 {object} kernal.ErrorResult "请求参数错误"
// @Failure 500 {object} kernal.ErrorResult "服务器内部错误"
// @Router /api/auction [post]
func (h *AuctionHandler) CreateAuction(ctx *gin.Context) {
	var req CreateAuctionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, kernal.NewErrorResult(1, err.Error()))
		return
	}
	// 创建拍卖品
	userID := ctx.Value("userID").(uint)
	cmd := application.CreateAuctionCommand{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		InitPrice:   req.InitPrice,
		Step:        req.StepPrice,
	}
	if err := h.auctionService.CreateAuction(&cmd); err != nil {
		ctx.JSON(http.StatusInternalServerError, kernal.NewErrorResult(1, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, kernal.NewDefaultSuccessResult())
}

// @Summary 获取最新的拍卖品
// @Description 获取最新的拍卖品
// @Tags 拍卖品
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} kernal.SuccessResult{data=application.AuctionBriefListDTO} "获取成功"
// @Failure 400 {object} kernal.ErrorResult "请求参数错误"
// @Failure 500 {object} kernal.ErrorResult "服务器内部错误"
// @Router /api/auction/latest [get]
func (h *AuctionHandler) ListLatestAuctions(ctx *gin.Context) {
	var req ListAuctionRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, kernal.NewErrorResult(1, err.Error()))
		return
	}
	// 如果没有分页参数，使用默认分页
	if req.Page == 0 || req.Size == 0 {
		req.Pagination = kernal.NewDefaultPagination()
	}
	// 获取最新的拍卖品
	auctions, err := h.auctionService.ListAuctions(&application.ListAuctionsQuery{
		Pagination: kernal.Pagination{
			Page: req.Page,
			Size: req.Size,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, kernal.NewErrorResult(1, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, auctions)
}

// @Summary 创建出价
// @Description 创建出价
// @Tags 出价
// @Accept json
// @Produce json
// @Param req body CreateBidRequest true "创建出价请求参数"
// @Success 200 {object} kernal.SuccessResult "创建成功"
// @Failure 400 {object} kernal.ErrorResult "请求参数错误"
// @Failure 500 {object} kernal.ErrorResult "服务器内部错误"
// @Router /api/auction/bid [post]
func (h *AuctionHandler) CreateBid(ctx *gin.Context) {
	var req CreateBidRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, kernal.NewErrorResult(1, err.Error()))
		return
	}
	// 创建出价
	userID := ctx.Value("userID").(uint)
	cmd := application.CreateBidCommand{
		AuctionID: req.AuctionID,
		UserID:    userID,
		Price:     req.Price,
	}
	if err := h.auctionService.CreateBid(&cmd); err != nil {
		ctx.JSON(http.StatusInternalServerError, kernal.NewErrorResult(1, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, kernal.NewDefaultSuccessResult())
}

// @Summary 获取拍卖品的最高出价
// @Description 获取拍卖品的最高出价
// @Tags 出价
// @Accept json
// @Produce json
// @Param req body GetHigestBidRequest true "获取拍卖品的最高出价请求参数"
// @Success 200 {object} kernal.SuccessResult{data=application.BidBriefListDTO} "获取成功"
// @Failure 400 {object} kernal.ErrorResult "请求参数错误"
// @Failure 500 {object} kernal.ErrorResult "服务器内部错误"
// @Router /api/auction/bid/higest [get]
func (h *AuctionHandler) GetHigestBid(ctx *gin.Context) {
	var req GetHigestBidRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, kernal.NewErrorResult(1, err.Error()))
		return
	}
	// 获取最高出价
	bids, err := h.auctionService.ListLatestBids(&application.ListLatestBidsQuery{
		Pagination: kernal.Pagination{
			Page: 1,
			Size: 1,
		},
		AuctionID: req.AuctionID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, kernal.NewErrorResult(1, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, kernal.NewSuccessResult(bids[0]))
}

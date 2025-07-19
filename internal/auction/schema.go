package auction

import "auctionsystem/internal/common"

type CreateRequestSchema struct {
	Title     string `json:"title" gorm:"column:title"`
	StartTime int64  `json:"start_time" gorm:"column:start_time"`
	EndTime   int64  `json:"end_time" gorm:"column:end_time"`
	InitPrice int64  `json:"init_price" gorm:"column:init_price"`
	Step      int64  `json:"step" gorm:"column:step"`
}

type CreateResponseSchema struct {
	common.ResponseSchema
}

type GetResponseSchema struct {
	common.ResponseSchema
}

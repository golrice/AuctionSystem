package auction

import "auctionsystem/internal/common"

type CreateRequestSchema struct {
	Title     string `json:"title" gorm:"column:title"`
	StartTime int64  `json:"start_time" gorm:"column:start_time"`
	EndTime   int64  `json:"end_time" gorm:"column:end_time"`
	InitPrice int64  `json:"init_price" gorm:"column:init_price"`
	Step      int64  `json:"step" gorm:"column:step"`
	UserID    uint   `json:"user_id" gorm:"column:user_id;"`
}

type CreateResponseSchema struct {
	common.ResponseSchema
}

type GetRequestSchema struct {
	ID uint `json:"id" gorm:"primarykey"`
}

type GetResponseSchema struct {
	common.ResponseSchema

	Data *Model `json:"data"`
}

type UpdateRequestSchema struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	Title     string `json:"title" gorm:"column:title"`
	StartTime int64  `json:"start_time" gorm:"column:start_time"`
	EndTime   int64  `json:"end_time" gorm:"column:end_time"`
	InitPrice int64  `json:"init_price" gorm:"column:init_price"`
	Step      int64  `json:"step" gorm:"column:step"`
	Status    int    `json:"status" gorm:"column:status"`
}

type UpdateResponseSchema struct {
	common.ResponseSchema
}

type DeleteRequestSchema struct {
	ID uint `json:"id" gorm:"primarykey"`
}

type DeleteResponseSchema struct {
	common.ResponseSchema
}

type ListRequestSchema struct {
	Status int `json:"status" gorm:"column:status"`
}

type ListResponseSchema struct {
	common.ResponseSchema

	Data []*Model `json:"data"`
}

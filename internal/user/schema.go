package user

import (
	"auctionsystem/internal/common"
	"time"
)

type GetRequestSchema struct {
	ID uint `json:"id" form:"id" query:"id" uri:"id"`
}

type GetResponseSchema struct {
	common.ResponseSchema

	CreatedAt time.Time `json:"created_at"`

	Name    string `json:"name" form:"name" query:"name" uri:"name"`
	Email   string `json:"email" form:"email" query:"email" uri:"email"`
	Balance int64  `json:"balance" form:"balance" query:"balance" uri:"balance"`
}

type UpdateRequestSchema struct {
	ID uint `json:"id" form:"id" query:"id" uri:"id"`
}

type UpdateResponseSchema struct {
	common.ResponseSchema
}

type DeleteRequestSchema struct {
	ID uint `json:"id" form:"id" query:"id" uri:"id"`
}

type DeleteResponseSchema struct {
	common.ResponseSchema
}

package user

import (
	"auctionsystem/internal/common"
	"context"

	"gorm.io/gorm"
)

func (s *Service) Get(request GetRequestSchema) (*GetResponseSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()

	response, err := s.repo.Get(ctx, &Model{Model: gorm.Model{ID: request.ID}})
	if err != nil {
		return nil, err
	}

	return &GetResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "success",
		},
		CreatedAt: response.CreatedAt,

		Name:    response.Name,
		Email:   response.Email,
		Balance: response.Balance,
	}, nil
}

func (s *Service) Update(request UpdateRequestSchema) (*UpdateResponseSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()

	if err := s.repo.Update(ctx, &Model{Model: gorm.Model{ID: request.ID}}); err != nil {
		return nil, err
	}

	return &UpdateResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "success",
		},
	}, nil
}

func (s *Service) Delete(request DeleteRequestSchema) (*DeleteResponseSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()

	if err := s.repo.Delete(ctx, &Model{Model: gorm.Model{ID: request.ID}}); err != nil {
		return nil, err
	}

	return &DeleteResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "success",
		},
	}, nil
}

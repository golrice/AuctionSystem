package auction

import (
	"auctionsystem/internal/common"
	"context"

	"gorm.io/gorm"
)

func (s *Service) Create(request *CreateRequestSchema) (*CreateResponseSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()

	model := &Model{
		Title:     request.Title,
		StartTime: request.StartTime,
		EndTime:   request.EndTime,
		InitPrice: request.InitPrice,
		Step:      request.Step,
		Status:    AuctionStatusWaiting,
	}
	if err := s.repo.Create(ctx, model); err != nil {
		return nil, err
	}
	return &CreateResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "success",
		},
	}, nil
}

func (s *Service) Get(request *GetRequestSchema) (*GetResponseSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()

	model := &Model{
		Model: gorm.Model{
			ID: request.ID,
		},
	}
	response, err := s.repo.Get(ctx, model)
	if err != nil {
		return nil, err
	}
	return &GetResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "success",
		},
		Data: response,
	}, nil
}

func (s *Service) Update(request *UpdateRequestSchema) (*UpdateResponseSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()

	model := &Model{
		Model: gorm.Model{
			ID: request.ID,
		},
		Title:     request.Title,
		StartTime: request.StartTime,
		EndTime:   request.EndTime,
		InitPrice: request.InitPrice,
		Step:      request.Step,
	}
	if err := s.repo.Update(ctx, model); err != nil {
		return nil, err
	}
	return &UpdateResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "success",
		},
	}, nil
}

func (s *Service) Delete(request *DeleteRequestSchema) (*DeleteResponseSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()

	model := &Model{
		Model: gorm.Model{
			ID: request.ID,
		},
	}
	if err := s.repo.Delete(ctx, model); err != nil {
		return nil, err
	}
	return &DeleteResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "success",
		},
	}, nil
}

func (s *Service) List(request *ListRequestSchema) (*ListResponseSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()

	model := &Model{}
	response, err := s.repo.List(ctx, model)
	if err != nil {
		return nil, err
	}
	return &ListResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "success",
		},
		Data: response,
	}, nil
}

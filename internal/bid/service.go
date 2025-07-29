package bid

import (
	"auctionsystem/internal/common"
	"context"
	"errors"
	"time"
)

func (s *Service) Create(request *CreateRequestSchema) (*CreateResponseSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.persistentRepo.Create(ctx, &Model{
		Description: Description{
			AuctionID: request.AuctionID,
			UserID:    request.UserID,
			Price:     request.Price,
		},
	}); err != nil {
		return nil, errors.New("create bid failed")
	}

	return &CreateResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "create bid success",
		},
	}, nil
}

func (s *Service) List(request *ListRequestSchema) (*ListResponseSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bids, err := s.persistentRepo.List(ctx, &Model{
		Description: Description{
			AuctionID: request.AuctionID,
		},
	})
	if err != nil {
		return nil, errors.New("list bid failed")
	}

	bidsDescription := make([]Description, 0, len(bids))
	for _, bid := range bids {
		bidsDescription = append(bidsDescription, bid.Description)
	}
	return &ListResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "list bid success",
		},
		Bids: bidsDescription,
	}, nil
}

func (s *Service) Delete(request *DeleteRequestSchema) (*DeleteResponseSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.persistentRepo.Delete(ctx, &Model{
		Description: Description{
			AuctionID: request.AuctionID,
		},
	}); err != nil {
		return nil, errors.New("delete bid failed")
	}

	return &DeleteResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "delete bid success",
		},
	}, nil
}

func (s *Service) Detail(request *DetailRequestSchema) (*DetailResponseSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bid, err := s.persistentRepo.Get(ctx, &Model{
		Description: Description{
			AuctionID: request.AuctionID,
		},
	})
	if err != nil {
		return nil, errors.New("get bid failed")
	}

	return &DetailResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "detail bid success",
		},
		Bid: bid.Description,
	}, nil
}

func (s *Service) Update(request *UpdateRequestSchema) (*UpdateResponseSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.persistentRepo.Update(ctx, &Model{
		Description: Description{
			AuctionID: request.AuctionID,
			Price:     request.Price,
		},
	}); err != nil {
		return nil, errors.New("update bid failed")
	}

	return &UpdateResponseSchema{
		ResponseSchema: common.ResponseSchema{
			Code: 0,
			Msg:  "update bid success",
		},
	}, nil
}

var _ BidService = &Service{}

package auction

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type AuctionService interface {
	Create(request *CreateRequestSchema) (*CreateResponseSchema, error)
	Get(request *GetRequestSchema) (*GetResponseSchema, error)
	Update(request *UpdateRequestSchema) (*UpdateResponseSchema, error)
	Delete(request *DeleteRequestSchema) (*DeleteResponseSchema, error)
	List(request *ListRequestSchema) (*ListResponseSchema, error)
}

type AuctionRepository interface {
	Create(ctx context.Context, model *Model) error
	Get(ctx context.Context, model *Model) (*Model, error)
	Update(ctx context.Context, model *Model) error
	Delete(ctx context.Context, model *Model) error
	List(ctx context.Context, model *Model) ([]*Model, error)
}

type Repository struct {
	db *gorm.DB
}

type Service struct {
	repo           AuctionRepository
	contextTimeout time.Duration
}

func NewAuctionRepository(db *gorm.DB) AuctionRepository {
	return &Repository{db: db}
}

func NewAuctionService(repo AuctionRepository, timeout time.Duration) AuctionService {
	return &Service{repo: repo, contextTimeout: timeout}
}

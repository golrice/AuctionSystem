package user

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, model *Model) error
	Get(ctx context.Context, model *Model) (*Model, error)
	Update(ctx context.Context, model *Model) error
	Delete(ctx context.Context, model *Model) error
}

type UserService interface {
	Get(request GetRequestSchema) (*GetResponseSchema, error)
	Update(request UpdateRequestSchema) (*UpdateResponseSchema, error)
	Delete(request DeleteRequestSchema) (*DeleteResponseSchema, error)
}

type Service struct {
	repo           UserRepository
	contextTimeout time.Duration
}

type Repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &Repository{db: db}
}

func NewUserService(repo UserRepository, contextTimeout time.Duration) UserService {
	return &Service{repo: repo, contextTimeout: contextTimeout}
}

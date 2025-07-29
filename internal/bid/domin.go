package bid

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BidRepository interface {
	Create(ctx context.Context, model *Model) error
	Get(ctx context.Context, model *Model) (*Model, error)
	Update(ctx context.Context, model *Model) error
	Delete(ctx context.Context, model *Model) error
	List(ctx context.Context, model *Model) ([]*Model, error)
}

type BidMemCache interface {
	Add(ctx context.Context, key string, score float64, member Description, expiration time.Duration) error
	List(ctx context.Context, key string) ([]Description, error)
}

type BidService interface {
	Create(request *CreateRequestSchema) (*CreateResponseSchema, error)
	Detail(request *DetailRequestSchema) (*DetailResponseSchema, error)
	Update(request *UpdateRequestSchema) (*UpdateResponseSchema, error)
	Delete(request *DeleteRequestSchema) (*DeleteResponseSchema, error)
	List(request *ListRequestSchema) (*ListResponseSchema, error)
}

type MemoryRepository struct {
	client *redis.Client
	mtx    sync.Mutex
}

type PersistentRepository struct {
	db *gorm.DB
}

type Service struct {
	cacheRepo      BidMemCache
	persistentRepo BidRepository
	contextTimeout time.Duration
}

func NewService(cacheRepo BidMemCache, persistentRepo BidRepository, contextTimeout time.Duration) *Service {
	return &Service{
		cacheRepo:      cacheRepo,
		persistentRepo: persistentRepo,
		contextTimeout: contextTimeout,
	}
}

func NewMemoryRepository(client *redis.Client) BidMemCache {
	return &MemoryRepository{
		client: client,
		mtx:    sync.Mutex{},
	}
}

func NewPersistentRepository(db *gorm.DB) BidRepository {
	return &PersistentRepository{
		db: db,
	}
}

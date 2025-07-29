package bid

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r *MemoryRepository) Add(ctx context.Context, key string, score float64, member Description, expiration time.Duration) error {
	jsonData, err := json.Marshal(member)
	if err != nil {
		return errors.New("add member failed")
	}

	r.client.ZAdd(ctx, key, redis.Z{
		Score:  score,
		Member: jsonData,
	})
	return nil
}

func (r *MemoryRepository) List(ctx context.Context, key string) ([]Description, error) {
	allMembersWithScores, err := r.client.ZRevRangeWithScores(ctx, "leaderboard", 0, -1).Result()
	if err != nil {
		return nil, errors.New("list leaderboard failed")
	}

	results := make([]Description, 0)
	for _, member := range allMembersWithScores {
		var description Description
		if err := json.Unmarshal([]byte(member.Member.(string)), &description); err != nil {
			return nil, errors.New("unmarshal leaderboard failed")
		}
		results = append(results, description)
	}
	return results, nil
}

func (r *PersistentRepository) Create(ctx context.Context, model *Model) error {
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return errors.New("create bid failed")
	}

	return nil
}

func (r *PersistentRepository) Get(ctx context.Context, model *Model) (*Model, error) {
	var result *Model
	if err := r.db.WithContext(ctx).First(&result).Error; err != nil {
		return nil, errors.New("get bid failed")
	}

	return result, nil
}

func (r *PersistentRepository) Update(ctx context.Context, model *Model) error {
	if err := r.db.WithContext(ctx).Updates(model).Error; err != nil {
		return errors.New("update bid failed")
	}
	return nil
}

func (r *PersistentRepository) Delete(ctx context.Context, model *Model) error {
	if err := r.db.WithContext(ctx).Delete(model).Error; err != nil {
		return errors.New("delete bid failed")
	}
	return nil
}

func (r *PersistentRepository) List(ctx context.Context, model *Model) ([]*Model, error) {
	var result []*Model
	if err := r.db.WithContext(ctx).Where(model).Find(&result).Error; err != nil {
		return nil, errors.New("list bid failed")
	}
	return result, nil
}

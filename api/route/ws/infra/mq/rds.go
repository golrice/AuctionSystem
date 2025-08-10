package mq

import (
	"auctionsystem/api/route/ws/domain"
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	redis *redis.Client
}

type RedisChannel struct {
	cancelFunc context.CancelFunc
	pubsub     *redis.PubSub
}

func NewRedisRepository(redis *redis.Client) *RedisRepository {
	return &RedisRepository{redis: redis}
}

func (r *RedisRepository) Subscribe(ctx context.Context, auctionID uint) domain.MQChannel {
	ctx, cancel := context.WithCancel(ctx)
	pubsub := r.redis.Subscribe(ctx, fmt.Sprintf("auction:%d", auctionID))
	return &RedisChannel{cancelFunc: cancel, pubsub: pubsub}
}

func (r *RedisRepository) Publish(ctx context.Context, msg domain.AuctionMessage) error {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return r.redis.Publish(ctx, fmt.Sprintf("auction:%d", msg.AuctionID), jsonMsg).Err()
}

func (c *RedisChannel) Receive() <-chan *redis.Message {
	return c.pubsub.Channel()
}

func (c *RedisChannel) Cancel() {
	c.cancelFunc()
	_ = c.pubsub.Close()
}

var _ domain.AuctionMessageRepository = (*RedisRepository)(nil)

var _ domain.MQChannel = (*RedisChannel)(nil)

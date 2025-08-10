package domain

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type AuctionMessageRepository interface {
	Subscribe(ctx context.Context, auctionID uint) MQChannel
	Publish(ctx context.Context, msg AuctionMessage) error
}

type MQChannel interface {
	Cancel()
	Receive() <-chan *redis.Message
}

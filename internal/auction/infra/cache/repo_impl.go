package cache

import (
	"auctionsystem/internal/auction/domain"
	"auctionsystem/pkg/config"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type AuctionCacheImpl struct {
	cache *redis.Client
}

func NewAuctionCacheImpl(cache *redis.Client) domain.AuctionCache {
	return &AuctionCacheImpl{cache: cache}
}

func (a *AuctionCacheImpl) CreateAuction(ctx context.Context, auction *domain.Auction) error {
	data, err := json.Marshal(auction)
	if err != nil {
		return err
	}
	cmd := a.cache.Set(ctx, fmt.Sprintf("auction:%d", auction.ID), data, config.RedisTTL*time.Second)
	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}

func (a *AuctionCacheImpl) FindAuctionByID(ctx context.Context, id uint) (*domain.Auction, error) {
	data, err := a.cache.Get(ctx, fmt.Sprintf("auction:%d", id)).Bytes()
	if err != nil {
		return nil, err
	}
	auction := &domain.Auction{}
	if err := json.Unmarshal(data, auction); err != nil {
		return nil, err
	}
	return auction, nil
}

func (a *AuctionCacheImpl) UpdateAuction(ctx context.Context, auction *domain.Auction) error {
	data, err := json.Marshal(auction)
	if err != nil {
		return err
	}
	cmd := a.cache.Set(ctx, fmt.Sprintf("auction:%d", auction.ID), data, config.RedisTTL*time.Second)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (a *AuctionCacheImpl) DeleteAuction(ctx context.Context, id uint) error {
	cmd := a.cache.Del(ctx, fmt.Sprintf("auction:%d", id))
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (a *AuctionCacheImpl) LoadAuctionLatestBid(ctx context.Context, auction *domain.Auction) (*domain.Bid, error) {
	cmd, err := a.cache.Get(ctx, fmt.Sprintf("auction:%d:bid", auction.ID)).Bytes()
	if err != nil {
		return nil, err
	}
	bid := &domain.Bid{}
	if err := json.Unmarshal(cmd, bid); err != nil {
		return nil, err
	}
	return bid, nil
}

func (a *AuctionCacheImpl) CreateBid(ctx context.Context, bid *domain.Bid) error {
	data, err := json.Marshal(bid)
	if err != nil {
		return err
	}
	cmd := a.cache.Set(ctx, fmt.Sprintf("auction:%d:bid", bid.AuctionID), data, config.RedisTTL*time.Second)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (a *AuctionCacheImpl) DeleteBid(ctx context.Context, auctionID uint) error {
	cmd := a.cache.Del(ctx, fmt.Sprintf("auction:%d:bid", auctionID))
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (a *AuctionCacheImpl) Lock(ctx context.Context, key string) error {
	return a.cache.SetNX(ctx, key, "1", config.RedisLockTTL*time.Second).Err()
}

func (a *AuctionCacheImpl) Unlock(ctx context.Context, key string) error {
	return a.cache.Del(ctx, key).Err()
}

var _ domain.AuctionCache = (*AuctionCacheImpl)(nil)

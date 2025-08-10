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

type AuctionRepositoryImpl struct {
	cache *redis.Client
}

func NewAuctionCacheImpl(cache *redis.Client) domain.AuctionCache {
	return &AuctionRepositoryImpl{cache: cache}
}

func (a *AuctionRepositoryImpl) CreateAuction(ctx context.Context, auction *domain.Auction) error {
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

func (a *AuctionRepositoryImpl) FindAuctionByID(ctx context.Context, id uint) (*domain.Auction, error) {
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

func (a *AuctionRepositoryImpl) UpdateAuction(ctx context.Context, auction *domain.Auction) error {
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

func (a *AuctionRepositoryImpl) DeleteAuction(ctx context.Context, id uint) error {
	cmd := a.cache.Del(ctx, fmt.Sprintf("auction:%d", id))
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (a *AuctionRepositoryImpl) LoadAuctionLatestBid(ctx context.Context, auction *domain.Auction) (*domain.Bid, error) {
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

func (a *AuctionRepositoryImpl) CreateBid(ctx context.Context, bid *domain.Bid) error {
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

var _ domain.AuctionCache = (*AuctionRepositoryImpl)(nil)

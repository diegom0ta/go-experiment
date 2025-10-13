package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"experiment/core/domain"
	"experiment/infra/cache"
)

const (
	// DefaultTTL sets cache expiration to 72 hours
	DefaultTTL = 72 * time.Hour
)

type WalletCache struct{}

func NewWalletCache() *WalletCache {
	return &WalletCache{}
}

// CacheWallet stores a wallet in Redis with default 72-hour TTL
func (wc *WalletCache) CacheWallet(ctx context.Context, wallet *domain.Wallet) error {
	return wc.CacheWalletWithTTL(ctx, wallet, DefaultTTL)
}

// CacheWalletWithTTL stores a wallet in Redis with custom TTL
func (wc *WalletCache) CacheWalletWithTTL(ctx context.Context, wallet *domain.Wallet, ttl time.Duration) error {
	data, err := json.Marshal(wallet)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("wallet:%s", wallet.ID)
	return cache.RedisClient.Set(ctx, key, data, ttl).Err()
}

// GetWallet retrieves a wallet from Redis
func (wc *WalletCache) GetWallet(ctx context.Context, walletID string) (*domain.Wallet, error) {
	key := fmt.Sprintf("wallet:%s", walletID)
	data, err := cache.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var wallet domain.Wallet
	err = json.Unmarshal([]byte(data), &wallet)
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}

// DeleteWallet removes a wallet from Redis
func (wc *WalletCache) DeleteWallet(ctx context.Context, walletID string) error {
	key := fmt.Sprintf("wallet:%s", walletID)
	return cache.RedisClient.Del(ctx, key).Err()
}

type OwnerCache struct{}

func NewOwnerCache() *OwnerCache {
	return &OwnerCache{}
}

// CacheOwner stores an owner in Redis with default 72-hour TTL
func (oc *OwnerCache) CacheOwner(ctx context.Context, owner *domain.Owner) error {
	return oc.CacheOwnerWithTTL(ctx, owner, DefaultTTL)
}

// CacheOwnerWithTTL stores an owner in Redis with custom TTL
func (oc *OwnerCache) CacheOwnerWithTTL(ctx context.Context, owner *domain.Owner, ttl time.Duration) error {
	data, err := json.Marshal(owner)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("owner:%s", owner.Email)
	return cache.RedisClient.Set(ctx, key, data, ttl).Err()
}

// GetOwner retrieves an owner from Redis
func (oc *OwnerCache) GetOwner(ctx context.Context, email string) (*domain.Owner, error) {
	key := fmt.Sprintf("owner:email:%s", email)
	data, err := cache.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var owner domain.Owner
	err = json.Unmarshal([]byte(data), &owner)
	if err != nil {
		return nil, err
	}

	return &owner, nil
}

// DeleteOwner removes an owner from Redis
func (oc *OwnerCache) DeleteOwner(ctx context.Context, email string) error {
	key := fmt.Sprintf("owner:email:%s", email)
	return cache.RedisClient.Del(ctx, key).Err()
}

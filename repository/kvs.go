package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tatuya-web/go-gin-template/config"
	"github.com/tatuya-web/go-gin-template/domain/model"
)

func NewKVS(ctx context.Context, cfg *config.Config) (*KVS, error) {
	cli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
	})
	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &KVS{Cli: cli}, nil
}

type KVS struct {
	Cli *redis.Client
}

func (k *KVS) Save(ctx context.Context, key string, userID model.UserID) error {
	id := int64(userID)
	return k.Cli.Set(ctx, key, id, 30*time.Minute).Err()
}

func (k *KVS) Load(ctx context.Context, key string) (model.UserID, error) {
	id, err := k.Cli.Get(ctx, key).Int64()
	if err != nil {
		return 0, fmt.Errorf("failed to get by %q: %w", key, ErrNotFound)
	}
	return model.UserID(id), nil
}

func (k *KVS) Delete(ctx context.Context, key string) error {
	if err := k.Cli.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete by %q: %w", key, ErrNotFound)
	}
	return nil
}

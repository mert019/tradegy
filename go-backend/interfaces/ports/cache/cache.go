package cache

import "context"

type ICache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expireSeconds int64) error
	Exists(ctx context.Context, key string) (bool, error)
}

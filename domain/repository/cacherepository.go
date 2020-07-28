package domain

import (
	"context"
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/ahmadrezamusthafa/assessment/pkg/cache"
)

type CacheRepositoryItf interface {
	GetObjCache(ctx context.Context, key string, destObj interface{}) error
	SetObjCache(ctx context.Context, key string, value interface{}, ttl int) error
	ForceExpireCache(ctx context.Context, keys ...string) (err error)
}

type CacheRepository struct {
	Redis *cache.AssessmentCache
}

func (repo CacheRepository) GetObjCache(ctx context.Context, key string, destObj interface{}) (err error) {
	err = errors.AddTrace(repo.Redis.GetObj(key, &destObj))
	return
}

func (repo CacheRepository) SetObjCache(ctx context.Context, key string, value interface{}, ttl int) (err error) {
	err = errors.AddTrace(repo.Redis.SetExObj(key, value, ttl))
	return
}

func (repo CacheRepository) ForceExpireCache(ctx context.Context, keys ...string) (err error) {
	err = errors.AddTrace(repo.Redis.ForceExpire(keys...))
	return
}

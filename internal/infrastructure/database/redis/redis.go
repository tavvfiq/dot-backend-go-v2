package redis

import (
	"context"
	"go-backend-v2/internal/domain"
	"log"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type CacheProvider interface {
	Store(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string, value interface{}) error
	Delete(ctx context.Context, key string) error
}

type Redis struct {
	client *redis.Client
	cache  *cache.Cache
}

func NewRedis(address string, password string, db int) CacheProvider {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db, // default db
	})
	cache := cache.New(&cache.Options{
		Redis:      client,
		LocalCache: cache.NewTinyLFU(1000, 5*time.Minute),
	})

	rd := &Redis{
		client: client,
		cache:  cache,
	}

	// check redis connection
	err := rd.Store(context.Background(), "test-connection", "")
	if err != nil {
		log.Panic(err)
	}
	log.Printf("successfully connect to redis: %v", address)
	return rd
}

func (r *Redis) Store(ctx context.Context, key string, value interface{}) error {
	err := r.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   5 * time.Minute,
	})
	if err != nil {
		log.Println(errors.Wrapf(err, "error storing %s", key))
		return domain.ErrInternalServer
	}
	return nil
}

func (r *Redis) Get(ctx context.Context, key string, value interface{}) error {
	err := r.cache.Get(ctx, key, value)
	if err != nil {
		log.Println(errors.Wrapf(err, "error retrieving %s", key))
		return domain.ErrInternalServer
	}
	return nil
}

func (r *Redis) Delete(ctx context.Context, key string) error {
	err := r.cache.Delete(ctx, key)
	if err != nil {
		log.Println(errors.Wrapf(err, "error deleting %s", key))
		return domain.ErrInternalServer
	}
	return nil
}

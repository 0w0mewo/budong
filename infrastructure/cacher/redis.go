package cacher

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var redisLogger *logrus.Entry

func init() {
	redisLogger = logrus.NewEntry(logrus.StandardLogger()).WithField("cacher", "redis")
}

// implement repo of shetu doamin
var _ KVStore = &RedisStore{}

type RedisStore struct {
	ctx context.Context
	rdb *redis.Client
}

func newRedisCache(ctx context.Context) *RedisStore {
	// TODO: use proper config
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", DB: 0})
	if err := client.Ping(ctx).Err(); err != nil {
		redisLogger.Panic(err)
	}

	return &RedisStore{
		ctx: ctx,
		rdb: client,
	}
}

func (rc *RedisStore) Add(key string, value []byte) error {
	return rc.rdb.Set(rc.ctx, key, value, time.Hour).Err()
}

func (rc *RedisStore) Get(key string, missDo func(key string) ([]byte, error)) ([]byte, error) {
	if rc.Exist(key) {
		redisLogger.Infof("cache hit: %s", key)
		return rc.rdb.Get(rc.ctx, key).Bytes()
	}

	b, err := missDo(key)
	if err != nil {
		return nil, err
	}
	rc.rdb.Set(rc.ctx, key, b, time.Hour)

	return b, nil
}

func (rc *RedisStore) Exist(key string) bool {
	return rc.rdb.Exists(rc.ctx, key).Val() != 0
}

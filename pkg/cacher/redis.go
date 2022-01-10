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

var _ KVStore = &RedisStore{}

type RedisStore struct {
	rdb     *redis.Client
	timeout time.Duration
}

func newRedisCache(addr string, timeout time.Duration) *RedisStore {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// TODO: use proper config
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", DB: 0})
	if err := client.Ping(ctx).Err(); err != nil {
		redisLogger.Panic(err)
	}

	return &RedisStore{
		rdb:     client,
		timeout: timeout,
	}
}

func (rc *RedisStore) Add(key string, value []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), rc.timeout)
	defer cancel()

	return rc.rdb.Set(ctx, key, value, time.Hour).Err()
}

func (rc *RedisStore) Get(key string, missDo func(key string) ([]byte, error)) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rc.timeout)
	defer cancel()

	if rc.Exist(key) {
		redisLogger.Infof("cache hit: %s", key)

		// renew TTL and get the key
		rc.rdb.Expire(ctx, key, time.Hour)
		return rc.rdb.Get(ctx, key).Bytes()
	}

	b, err := missDo(key)
	if err != nil {
		return nil, err
	}
	rc.rdb.Set(ctx, key, b, time.Hour)

	return b, nil
}

func (rc *RedisStore) Exist(key string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), rc.timeout)
	defer cancel()

	return rc.rdb.Exists(ctx, key).Val() != 0
}

func (rc *RedisStore) Close() error {
	return rc.rdb.Close()
}

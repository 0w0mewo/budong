package cacher

import "time"

type StoreType int

const (
	REDIS = iota
	MEM
)

// key-value based storage
type KVStore interface {
	Add(key string, value []byte) error
	Get(key string, missDo func(key string) ([]byte, error)) ([]byte, error)
	Exist(key string) bool
}

// redis based presistence storage
func NewRedisCache(addr string) KVStore {
	return newRedisCache(addr, 3*time.Second)
}

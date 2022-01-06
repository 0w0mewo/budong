package cacher

import (
	"fmt"
	"sync"
)

var _ KVStore = &InMemStore{}

type InMemStore struct {
	mem    map[string][]byte
	rwlock *sync.RWMutex
}

func NewInMemStore() *InMemStore {
	return &InMemStore{
		mem:    make(map[string][]byte),
		rwlock: &sync.RWMutex{},
	}
}

func (im *InMemStore) Get(key string, missDo func(key string) ([]byte, error)) ([]byte, error) {
	im.rwlock.RLock()
	if im.Exist(key) {
		im.rwlock.RUnlock()
		return im.mem[key], nil
	}

	b, err := missDo(key)
	if err != nil {
		return nil, err
	}

	im.rwlock.Lock()
	im.mem[key] = b
	im.rwlock.Unlock()

	return b, err
}

func (im *InMemStore) Add(key string, value []byte) error {
	im.rwlock.Lock()
	defer im.rwlock.Unlock()

	if exist := im.Exist(key); exist {
		return fmt.Errorf("%s exist", key)
	}

	im.mem[key] = value

	return nil
}

func (im *InMemStore) Exist(key string) bool {
	_, exist := im.mem[key]

	return exist
}

func (im *InMemStore) Close() error {
	return nil
}

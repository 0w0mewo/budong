package setu

import "errors"

var (
	ErrNotExistInDB    = errors.New("not exist in db")
	ErrNotExistInCache = errors.New("not exist in cache")
	ErrCacheHit        = errors.New("cache hit")
	ErrExistInDB       = errors.New("already in DB")
)

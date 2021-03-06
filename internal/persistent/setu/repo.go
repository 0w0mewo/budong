package setu

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/0w0mewo/budong/internal/persistent"
	"github.com/0w0mewo/budong/pkg/cacher"
	"github.com/0w0mewo/budong/pkg/domain/shetu"

	"github.com/sirupsen/logrus"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var _ shetu.Repo = &SetuRepo{}

type selectorFn func(v interface{}) ([]byte, error)

type setuRepoProvider interface {
	Create(setu *shetu.SetuInfo) (*shetu.Setu, error)
	SelectById(id int) ([]byte, error)
	SelectByTitle(title string) ([]byte, error)
	GetAmount() int64
	ListInventory(offset int64, limit int64) ([]*shetu.SetuInfo, error)
	SelectRandomId() (int, error)
	IsInDB(id int) bool
	Close() error
}

type SetuRepo struct {
	cache  cacher.KVStore
	db     setuRepoProvider
	logger *logrus.Entry
}

func NewSetuRepo(cacheProvider cacher.StoreType, repoProvider persistent.RepoProvider, dsn string, cacheDsn string) *SetuRepo {
	var cache cacher.KVStore
	var db setuRepoProvider

	switch cacheProvider {
	case cacher.REDIS:
		cache = cacher.NewRedisCache(cacheDsn)
	case cacher.MEM:
		cache = cacher.NewInMemStore()
	}

	switch repoProvider {
	case persistent.SQLITE:
		db = newSetuSqlDB(dsn)
	case persistent.MONGO:
		db = newSetuMongoDB(dsn)
	}

	return &SetuRepo{
		cache:  cache,
		db:     db,
		logger: logrus.StandardLogger().WithField("module", "setu_repo"),
	}
}

// get setu image bytes from redis by id.
// fetch it from database if it's not exist in the redis cache.
func (sr *SetuRepo) GetById(id int) ([]byte, error) {
	return sr.GetBy(id, strconv.Itoa(id), sr.selectSetuById)
}

// get setu image bytes from redis by title.
// fetch it from database if it's not exist in the redis cache.
func (sr *SetuRepo) GetByTitle(title string) ([]byte, error) {
	return sr.GetBy(title, title, sr.selectSetuByTitle)
}

// add setu
func (sr *SetuRepo) AddSetu(setu *shetu.SetuInfo) error {
	// setu is cached
	if sr.existCache(setu.Id) {
		sr.logger.Infof("cache hit when add setu: %d", setu.Id)
		return ErrCacheHit
	}

	// or in DB
	if sr.db.IsInDB(setu.Id) {
		sr.logger.Infof("db hit when add setu: %d", setu.Id)
		return ErrExistInDB
	}

	sr.logger.Infof("adding %d to DB", setu.Id)

	newRow, err := sr.db.Create(setu)
	if err != nil {
		return err
	}

	sr.logger.Infof("added %d to DB", setu.Id)

	// cache the image bytes
	return sr.cache.Add(setu.Key(), newRow.Data)

}

// list all
func (sr *SetuRepo) PaginatedInventory(page int64, pageLimit int64) ([]*shetu.SetuInfo, error) {
	return sr.db.ListInventory(page, pageLimit)
}

func (sr *SetuRepo) Count() int64 {
	return sr.db.GetAmount()
}

func (sr *SetuRepo) GetBy(val interface{}, cacheKey string, selector selectorFn) ([]byte, error) {
	// if the image is not in redis, try to fetch it from database
	// and cache the fetched image
	fetchFromDB := func(key string) ([]byte, error) {
		b, err := selector(val)
		if err != nil {
			return nil, err
		}

		sr.cache.Add(key, b)
		return b, nil
	}

	res, err := sr.cache.Get(cacheKey, fetchFromDB)
	if err != nil {
		return nil, err
	}

	// the image is cached, return it
	return res, nil
}

// get setu image bytes from DB by id, return error if not found
func (sr *SetuRepo) selectSetuById(id interface{}) ([]byte, error) {
	return sr.db.SelectById(id.(int))

}

// get setu image bytes from DB by title, return error if not found
func (sr *SetuRepo) selectSetuByTitle(title interface{}) ([]byte, error) {
	return sr.db.SelectByTitle(title.(string))

}

// whether a given setu id exist in cache
func (sr *SetuRepo) existCache(id int) bool {
	return sr.cache.Exist(strconv.Itoa(id))
}

func (sr *SetuRepo) Random() (int, error) {
	return sr.db.SelectRandomId()
}

func (sr *SetuRepo) Close() error {
	err := sr.db.Close()
	if err != nil {
		return err
	}

	return sr.cache.Close()
}

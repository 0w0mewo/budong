package persistent

import (
	"context"
	"strconv"

	"github.com/0w0mewo/budong/config"
	"github.com/0w0mewo/budong/domain/shetu"
	"github.com/0w0mewo/budong/infrastructure/cacher"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _ shetu.Repo = &SetuRepo{}

type selectorFn func(v interface{}) ([]byte, error)

type SetuRepo struct {
	cache  cacher.KVStore
	db     *gorm.DB
	logger *logrus.Entry
}

func NewSetuRepo(ctx context.Context, t cacher.StoreType, dsn string) *SetuRepo {
	var cache cacher.KVStore

	switch t {
	case cacher.REDIS:
		cache = cacher.NewRedisCache(ctx, config.GlobalConfig.RedisAddr())
	case cacher.MEM:
		cache = cacher.NewInMemStore()
	default:
		cache = cacher.NewRedisCache(ctx, config.GlobalConfig.RedisAddr())
	}

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	db = db.WithContext(ctx)

	db.AutoMigrate(&shetu.Setu{})

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
	if sr.existDB(setu.Id) {
		sr.logger.Infof("db hit when add setu: %d", setu.Id)
		return ErrExistInDB
	}

	sr.logger.Infof("adding %d to DB", setu.Id)

	// fetch it
	newRow, err := shetu.SetuFromSetuInfo(setu, true)
	if err != nil {
		return err
	}

	// add to DB
	err = sr.db.Model(&shetu.Setu{Id: setu.Id}).Create(newRow).Error
	if err != nil {
		return err
	}

	sr.logger.Infof("added %d to DB", setu.Id)

	// cache the image bytes
	return sr.cache.Add(setu.Key(), newRow.Data)

}

// list all
func (sr *SetuRepo) PaginatedInventory(page uint64, pageLimit uint64) ([]*shetu.SetuInfo, error) {
	var dbres []shetu.Setu

	offset := (page - 1) * pageLimit
	err := sr.db.Select("id", "title", "url", "uid", "is_r18").
		Offset(int(offset)).Limit(int(pageLimit)).
		Find(&dbres).Error
	if err != nil {
		return nil, err
	}

	res := make([]*shetu.SetuInfo, 0)
	for _, dbr := range dbres {
		res = append(res, shetu.SetuToSetuInfo(&dbr))
	}

	return res, err

}

func (sr *SetuRepo) Count() uint64 {
	var cnt int64
	sr.db.Model(&shetu.Setu{}).Count(&cnt)

	return uint64(cnt)
}

func (sr *SetuRepo) GetBy(val interface{}, cacheKey string, selector selectorFn) ([]byte, error) {
	// if the image is not in redis, try to fetch it from database
	// and cache the fetched image
	fetchFromDB := func(key string) ([]byte, error) {
		b, err := selector(val)
		if err != nil {
			return nil, err
		}

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
	cond := &shetu.Setu{Id: id.(int)}
	table := sr.db.Model(cond)
	res := table.First(cond)

	if res.RowsAffected == 0 {
		return nil, ErrNotExistInDB
	}

	return cond.Data, nil
}

// get setu image bytes from DB by title, return error if not found
func (sr *SetuRepo) selectSetuByTitle(title interface{}) ([]byte, error) {
	cond := &shetu.Setu{Title: title.(string)}
	table := sr.db.Model(cond)
	res := table.First(cond)

	if res.RowsAffected == 0 {
		return nil, ErrNotExistInDB
	}

	return cond.Data, nil
}

// whether a given setu id exist in cache
func (sr *SetuRepo) existCache(id int) bool {
	return sr.cache.Exist(strconv.Itoa(id))
}

// whether a given setu id exist in DB
func (sr *SetuRepo) existDB(id int) bool {
	cond := &shetu.Setu{Id: id}
	table := sr.db.Model(cond)
	res := table.First(cond)

	return res.RowsAffected != 0
}

func (sr *SetuRepo) Random() (int, error) {
	res := &shetu.Setu{}
	err := sr.db.Model(res).Select("id").Order("random()").First(res).Error

	return res.Id, err
}

package service

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/0w0mewo/budong/internal/infrastructure/cacher"
	"github.com/0w0mewo/budong/internal/infrastructure/persistent"
	"github.com/0w0mewo/budong/pkg/domain/shetu"
	"github.com/0w0mewo/budong/pkg/domain/upstream"

	"github.com/sirupsen/logrus"
)

type Service interface {
	RequestSetu(num int, isR18 bool) error
	GetSetuFromDB(id int) ([]byte, error)
	GetInventory(page, pageLimit uint64) ([]*shetu.SetuInfo, error)
	RandomSetu() ([]byte, error)
	Count() uint64
	Shutdown()
}

// implement service
var _ Service = &SetuService{}

type SetuService struct {
	wg     *sync.WaitGroup
	store  shetu.Repo // abstract storage
	logger *logrus.Entry
}

func NewSetuService(dsn string) *SetuService {
	db := persistent.NewSetuRepo(cacher.REDIS, dsn)
	return &SetuService{
		store:  db,
		wg:     &sync.WaitGroup{},
		logger: logrus.StandardLogger().WithField("module", "setu server"),
	}
}

// fetch setu from upstream
func (ss *SetuService) RequestSetu(num int, isR18 bool) error {
	return ss.fetchSetu(num, isR18, "")

}

// randomly select setu
func (ss *SetuService) RandomSetu() ([]byte, error) {
	id, err := ss.store.Random()
	if err != nil {
		return nil, err
	}

	return ss.store.GetById(id)
}

// get setu image bytes by id
func (ss *SetuService) GetSetuFromDB(id int) ([]byte, error) {
	b, err := ss.store.GetById(id)
	if err != nil {
		ss.logger.Errorln(err)
		return nil, err
	}

	return b, err
}

// setu inventory info
func (ss *SetuService) GetInventory(page, pageLimit uint64) ([]*shetu.SetuInfo, error) {
	setus, err := ss.store.PaginatedInventory(page, pageLimit)
	if err != nil {
		ss.logger.Errorln(err)
		return nil, err
	}

	return setus, nil
}

// number of setu in inventory
func (ss *SetuService) Count() uint64 {
	return uint64(ss.store.Count())
}

func (ss *SetuService) Shutdown() {
	ss.store.Close()
	ss.wg.Wait()

	ss.logger.Infoln("common service shutdown")
}

// request setu, fetch and store them into repo
func (ss *SetuService) fetchSetu(num int, r18 bool, keyword string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	setu, err := upstream.ReqSetuWithOption(ctx, http.DefaultClient, &upstream.Options{
		Num:     num,
		IsR18:   r18,
		Keyword: keyword,
	})
	if err != nil {
		ss.logger.Errorln(err)
		return err
	}

	ss.logger.Info("requested setu")

	for _, s := range setu.Data {
		ss.wg.Add(1)
		go func(s *shetu.SetuInfo) {
			defer ss.wg.Done()

			err := ss.store.AddSetu(s)
			if err != nil {
				if err == persistent.ErrCacheHit || err == persistent.ErrExistInDB {
					return
				}
				ss.logger.Errorln("add setu:", err)
				return
			}
		}(s)

	}

	ss.wg.Wait()

	return nil

}

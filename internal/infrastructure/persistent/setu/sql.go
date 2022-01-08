package setu

import (
	"github.com/0w0mewo/budong/pkg/domain/shetu"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _ setuRepoProvider = &setuSqlDB{}

type setuSqlDB struct {
	db     *gorm.DB
	logger *logrus.Entry
}

func newSetuSqlDB(dsn string) *setuSqlDB {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&shetu.Setu{})

	return &setuSqlDB{
		db:     db,
		logger: logrus.StandardLogger().WithField("module", "setu_repo"),
	}
}

func (s *setuSqlDB) Create(setu *shetu.SetuInfo) (*shetu.Setu, error) {
	// fetch it
	newRow, err := shetu.SetuFromSetuInfo(setu, true)
	if err != nil {
		return nil, err
	}

	// add to DB
	err = s.db.Model(&shetu.Setu{Id: setu.Id}).Create(newRow).Error
	if err != nil {
		return nil, err
	}

	return newRow, nil
}

func (s *setuSqlDB) SelectById(id int) ([]byte, error) {
	cond := &shetu.Setu{Id: id}
	table := s.db.Model(cond)
	res := table.First(cond)

	if res.RowsAffected == 0 {
		return nil, ErrNotExistInDB
	}

	return cond.Data, nil
}

func (s *setuSqlDB) SelectByTitle(title string) ([]byte, error) {
	cond := &shetu.Setu{Title: title}
	table := s.db.Model(cond)
	res := table.First(cond)

	if res.RowsAffected == 0 {
		return nil, ErrNotExistInDB
	}

	return cond.Data, nil
}

func (s *setuSqlDB) ListInventory(page uint64, pageLimit uint64) ([]*shetu.SetuInfo, error) {
	var dbres []shetu.Setu

	offset := (page - 1) * pageLimit
	err := s.db.Select("id", "title", "url", "uid", "is_r18").
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

func (s *setuSqlDB) GetAmount() uint64 {
	var cnt int64
	s.db.Model(&shetu.Setu{}).Count(&cnt)

	return uint64(cnt)
}

func (s *setuSqlDB) SelectRandomId() (int, error) {
	res := &shetu.Setu{}
	err := s.db.Model(res).Select("id").Order("random()").First(res).Error

	return res.Id, err
}

func (s *setuSqlDB) Close() error {
	return s.db.Commit().Error
}

// whether a given setu id exist in DB
func (sr *setuSqlDB) IsInDB(id int) bool {
	cond := &shetu.Setu{Id: id}
	table := sr.db.Model(cond)
	res := table.First(cond)

	return res.RowsAffected != 0
}

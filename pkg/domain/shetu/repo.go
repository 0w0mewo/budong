package shetu

// repo of setu domain
// exposed repo interface to infstrature layer
type Repo interface {
	GetById(id int) ([]byte, error)
	GetByTitle(title string) ([]byte, error)
	AddSetu(setu *SetuInfo) error
	PaginatedInventory(page int64, pageLimit int64) ([]*SetuInfo, error)
	Count() int64
	Random() (int, error)
	Close() error
}

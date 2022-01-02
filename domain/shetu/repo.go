package shetu

// repo of setu domain
// exposed repo interface to infstrature layer
type Repo interface {
	GetById(id int) ([]byte, error)
	GetByTitle(title string) ([]byte, error)
	GetImgTypeById(id int) (string, error)
	AddSetu(setu *SetuInfo) error
	PaginatedInventory(page uint64, pageLimit uint64) ([]*SetuInfo, error)
	Count() uint64
}

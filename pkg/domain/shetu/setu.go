package shetu

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/0w0mewo/budong/utils"
)

// setu data model in DB view
type Setu struct {
	Id    int `gorm:"primaryKey"`
	Title string
	Url   string
	Uid   int
	IsR18 bool   `bson:"is_r18"`
	Data  []byte `bson:"data" gorm:"index"`
}
type SetuInfo struct {
	Id    int               `json:"pid"`
	Title string            `json:"title"`
	Uid   int               `json:"uid"`
	Url   map[string]string `json:"urls"`
	IsR18 bool              `json:"r18"`
}

func (s *SetuInfo) String() string {
	return fmt.Sprintf("id: %d, title: %s, url: %s, uid: %d, is r18?: %v\n",
		s.Id, s.Title, s.URL(), s.Uid, s.IsR18)
}

// cache key
func (s *SetuInfo) Key() string {
	return strconv.Itoa(s.Id)
}

// url where the image pointed to
func (s *SetuInfo) URL() string {
	return s.Url["original"]
}

func (s *SetuInfo) FetchWithHook(retries int, fn func() error) ([]byte, error) {
	var err error
	var b []byte
	err = utils.Retry(retries, func() error {
		b, err = utils.HttpGetBytes(http.DefaultClient, s.URL())
		if err != nil {
			return err
		}

		// execute hooker after success
		if fn != nil {
			err = fn()
			if err != nil {
				return err
			}
		}

		return nil
	})

	return b, err
}

// download the image bytes from corrosponding url
func (s *SetuInfo) Fetch(retries int) ([]byte, error) {
	return s.FetchWithHook(retries, func() error {
		return nil
	})
}

// setu DB record to standard setu meta info
func SetuToSetuInfo(setu *Setu) *SetuInfo {
	res := &SetuInfo{
		Id:    setu.Id,
		Title: setu.Title,
		Uid:   setu.Uid,
		Url:   make(map[string]string),
		IsR18: setu.IsR18,
	}

	res.Url["original"] = setu.Url

	return res
}

// standard setu meta info to setu DB record, download the image if needed
func SetuFromSetuInfo(setu *SetuInfo, needFetch bool) (*Setu, error) {

	sdbr := &Setu{
		Id:    setu.Id,
		Url:   setu.URL(),
		Title: setu.Title,
		Uid:   setu.Uid,
		IsR18: setu.IsR18,
	}

	if needFetch {
		b, err := setu.Fetch(5)
		if err != nil {
			return sdbr, err
		}

		sdbr.Data = make([]byte, 0)
		sdbr.Data = append(sdbr.Data, b...)

	}

	return sdbr, nil
}

package upstream

import (
	"fmt"
	"strings"

	"github.com/0w0mewo/budong/domain/shetu"
)

// setu api request
type req struct {
	IsR18   int    `json:"r18"`
	Keyword string `json:"keyword"`
	Num     int    `json:"num"`
}

// setu api response
type SetuUpstreamResp struct {
	ErrMsg string            `json:"error"`
	Data   []*shetu.SetuInfo `json:"data"`
}

func (si *SetuUpstreamResp) String() string {
	setuTemp := `Title: %s, Url: %s, Is R18? : %v`

	setuBody := &strings.Builder{}

	for _, setu := range si.Data {
		setuBody.WriteString(fmt.Sprintf(setuTemp, setu.Title, setu.URL(), setu.IsR18))
	}

	return setuBody.String()
}

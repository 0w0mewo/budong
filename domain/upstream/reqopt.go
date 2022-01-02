package upstream

import (
	"context"
	"net/http"
	"strconv"
)

const url = "https://api.lolicon.app/setu/v2"

// setu option
type Options struct {
	Num     int
	IsR18   bool
	Keyword string
}

func newSetuReq(ctx context.Context, opt *Options) (*http.Request, error) {
	req := &req{
		IsR18:   0,
		Keyword: opt.Keyword,
		Num:     opt.Num,
	}

	// parse r18 option
	if opt.IsR18 {
		req.IsR18 = 1
	} else {
		req.IsR18 = 0
	}

	ret, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	// build query and add it to the request url
	parmas := ret.URL.Query()
	parmas.Add("r18", strconv.Itoa(req.IsR18))
	parmas.Add("keyword", req.Keyword)
	parmas.Add("num", strconv.Itoa(req.Num))
	ret.URL.RawQuery = parmas.Encode()

	return ret, nil

}

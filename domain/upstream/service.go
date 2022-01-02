package upstream

import (
	"context"
	"encoding/json"
	"net/http"
)

// request one setu
func ReqSetu(ctx context.Context, client *http.Client) (*SetuUpstreamResp, error) {
	return ReqSetuWithOption(ctx, client, &Options{
		Num:     1,
		IsSetu:  true,
		Keyword: "",
	})
}

// send request to setu api with option
func ReqSetuWithOption(ctx context.Context, client *http.Client, reqOpt *Options) (*SetuUpstreamResp, error) {

	// new setu request
	req, err := newSetuReq(ctx, reqOpt)
	if req == nil {
		return nil, err
	}

	// send request via client
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// decode json into setu api response
	result := &SetuUpstreamResp{}
	err = json.NewDecoder(res.Body).Decode(result)
	if err != nil {
		return nil, err
	}

	return result, nil

}

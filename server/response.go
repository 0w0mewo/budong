package server

import "github.com/0w0mewo/budong/domain/shetu"

type Resp struct {
	ErrMsg string            `json:"error"`
	Infos  []*shetu.SetuInfo `json:"inventory"`
	Count  int               `json:"count"`
}

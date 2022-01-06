package grpc

import (
	"context"

	"github.com/0w0mewo/budong/domain/shetu"
	"github.com/0w0mewo/budong/server"
	setupb "github.com/0w0mewo/budong/server/grpc/setu"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ setupb.SetuServiceServer = &GrpcServer{}

type GrpcServer struct {
	serve server.Service
	setupb.UnimplementedSetuServiceServer
}

// fetch setu and store to DB
func (gs *GrpcServer) Fetch(ctx context.Context,
	req *setupb.FetchReq) (*setupb.ErrResp, error) {

	err := gs.serve.RequestSetu(int(req.Amount), false) // 不可以色色
	if err != nil {
		emsg := err.Error()
		return &setupb.ErrResp{ErrMsg: emsg},
			status.Error(codes.Internal, emsg)
	}

	return &setupb.ErrResp{ErrMsg: "ok"}, nil
}

// get setu inventory
func (gs *GrpcServer) GetInventory(ctx context.Context,
	req *setupb.InventoryReq) (*setupb.InventoryResp, error) {

	// get page param
	page, size := req.Page, req.PageLimit

	// ensure page size is between 0 and 100
	if size > 50 || size < 1 {
		emsg := ErrPageSize.Error()
		return &setupb.InventoryResp{
			Err: &setupb.ErrResp{ErrMsg: emsg},
		}, status.Error(codes.InvalidArgument, emsg)
	}

	// ensure page is in valid rangeerr.Error()
	if page < 1 || page > gs.serve.Count()/size+1 {
		emsg := ErrPageRange.Error()
		return &setupb.InventoryResp{
			Err: &setupb.ErrResp{ErrMsg: emsg},
		}, status.Error(codes.InvalidArgument, emsg)
	}

	data, err := gs.serve.GetInventory(page, size)
	if err != nil {
		emsg := err.Error()
		return nil, status.Error(codes.Internal, emsg)
	}

	ret := SetuInfosToInventory(data)

	return ret, nil

}

func SetuInfosToInventory(setus []*shetu.SetuInfo) *setupb.InventoryResp {
	ret := &setupb.InventoryResp{
		Err:  &setupb.ErrResp{ErrMsg: "ok"},
		Info: make([]*setupb.InventoryResp_SetuInfo, 0),
	}

	for _, s := range setus {
		ret.Info = append(ret.Info,
			&setupb.InventoryResp_SetuInfo{
				Id:    int64(s.Id),
				Title: s.Title,
				Uid:   int64(s.Uid),
				Url:   s.URL(),
				IsR18: s.IsR18,
			})

	}
	return ret
}

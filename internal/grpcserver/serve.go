package grpcserver

import (
	"context"
	"io"
	"net"

	"github.com/0w0mewo/budong/pkg/domain/shetu"
	"github.com/0w0mewo/budong/pkg/service"
	"github.com/0w0mewo/budong/pkg/setupb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ setupb.SetuServiceServer = &SetuGrpcServer{}

type SetuGrpcServer struct {
	serve service.Service
	setupb.UnimplementedSetuServiceServer
	listener net.Listener
	server   *grpc.Server

	logger    *logrus.Entry
	running   bool
	chunkSize int
}

func NewSetuGrpcServer(addr, dsn string) *SetuGrpcServer {
	logger := logrus.StandardLogger().WithField("module", "grpc server")

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	gserver := grpc.NewServer()

	ret := &SetuGrpcServer{
		logger:    logger,
		serve:     service.NewSetuService(dsn),
		listener:  listener,
		server:    gserver,
		chunkSize: 32 * 1024, // 32KiB
	}

	setupb.RegisterSetuServiceServer(gserver, ret)

	return ret

}

// run server
func (sgs *SetuGrpcServer) Run() {
	if !sgs.running {
		sgs.running = true

		sgs.logger.Infoln("setu grpc running")
		err := sgs.server.Serve(sgs.listener)
		if err != nil {
			sgs.logger.Warnln(err)
		}

	}

}

// shutdown server
func (sgs *SetuGrpcServer) Shutdown() {
	sgs.serve.Shutdown()
	sgs.server.GracefulStop()
	sgs.running = false

	sgs.logger.Infoln("setu grpc server shutdown")
}

// fetch setu and store to DB
func (sgs *SetuGrpcServer) Fetch(ctx context.Context,
	req *setupb.FetchReq) (*setupb.FetchResp, error) {

	err := sgs.serve.RequestSetu(int(req.Amount), req.R18)
	if err != nil {
		emsg := err.Error()
		return &setupb.FetchResp{ErrMsg: emsg},
			status.Error(codes.Internal, emsg)
	}

	return &setupb.FetchResp{ErrMsg: "ok"}, nil
}

// get setu inventory
func (sgs *SetuGrpcServer) GetInventory(ctx context.Context,
	req *setupb.InventoryReq) (*setupb.InventoryResp, error) {

	// get page param
	page, size := req.Page, req.PageLimit

	// ensure page size is between 0 and 100
	if size > 50 || size < 1 {
		emsg := ErrPageSize.Error()
		return &setupb.InventoryResp{
			Err: &setupb.FetchResp{ErrMsg: emsg},
		}, status.Error(codes.InvalidArgument, emsg)
	}

	// ensure page is in valid rangeerr.Error()
	if page < 1 || page > sgs.serve.Count()/size+1 {
		emsg := ErrPageRange.Error()
		return &setupb.InventoryResp{
			Err: &setupb.FetchResp{ErrMsg: emsg},
		}, status.Error(codes.InvalidArgument, emsg)
	}

	data, err := sgs.serve.GetInventory(page, size)
	if err != nil {
		emsg := err.Error()
		return nil, status.Error(codes.Internal, emsg)
	}

	ret := setuInfosToInventory(data)

	return ret, nil

}

func (sgs *SetuGrpcServer) GetSetuById(req *setupb.SetuReq, stream setupb.SetuService_GetSetuByIdServer) error {
	img, err := sgs.serve.GetSetuFromDB(int(req.Id))
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return writeToStream(img, stream, sgs.chunkSize)
}

func (sgs *SetuGrpcServer) Random(req *setupb.RandomReq, stream setupb.SetuService_RandomServer) error {
	imgId, err := sgs.serve.RandomSetu()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return sgs.GetSetuById(&setupb.SetuReq{Id: int64(imgId)}, stream) // service is same as GetSetuById
}

// convert SetuInfo slice to gRPC InventoryResp
func setuInfosToInventory(setus []*shetu.SetuInfo) *setupb.InventoryResp {
	ret := &setupb.InventoryResp{
		Err:  &setupb.FetchResp{ErrMsg: "ok"},
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

// write to stream service until EOF
func writeToStream(src io.Reader, dst setupb.SetuService_GetSetuByIdServer, chunkSize int) error {
	buf := make([]byte, chunkSize)
	for {
		n, err := src.Read(buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}

			return err
		}

		dst.Send(&setupb.SetuResp{
			Chunk: buf[:n],
		})
	}

}

package main

import (
	"sync"

	"github.com/0w0mewo/budong/config"
	"github.com/0w0mewo/budong/internal/service/grpcserver"
	"github.com/0w0mewo/budong/utils"
)

func main() {
	cfg := config.LoadConfig()

	var wg sync.WaitGroup

	inst := grpcserver.NewSetuGrpcServer(cfg.ServiceAddr(), cfg.DSN(), cfg.RedisAddress())

	wg.Add(1)
	go func() {
		defer wg.Done()
		inst.Run()
	}()

	<-utils.WaitForSignal()

	inst.Shutdown()
	wg.Wait()
}

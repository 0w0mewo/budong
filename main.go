package main

import (
	"sync"

	"github.com/0w0mewo/budong/config"
	"github.com/0w0mewo/budong/internal/grpcserver"
	"github.com/0w0mewo/budong/utils"
)

func main() {
	var wg sync.WaitGroup

	inst := grpcserver.NewSetuGrpcServer(config.GlobalConfig.Addr(), config.GlobalConfig.DB())

	wg.Add(1)
	go func() {
		defer wg.Done()
		inst.Run()
	}()

	<-utils.WaitForSignal()

	inst.Shutdown()
	wg.Wait()
}

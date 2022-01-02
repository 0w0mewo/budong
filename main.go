package main

import (
	"context"

	"github.com/0w0mewo/budong/server"
	"github.com/0w0mewo/budong/utils"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	serve := server.NewSetuService(ctx, "/tmp/test.db")

	apiserver := server.NewRestfulServer(":9999", serve)
	apiserver.Init()
	go apiserver.Run()

	<-utils.WaitForSignal()

	cancel()
	apiserver.Shutdown()

}

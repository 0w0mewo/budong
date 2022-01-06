package main

import (
	"context"

	"github.com/0w0mewo/budong/config"
	"github.com/0w0mewo/budong/server"
	"github.com/0w0mewo/budong/server/httpserver"
	"github.com/0w0mewo/budong/utils"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	serve := server.NewSetuService(ctx, config.GlobalConfig.DB())

	apiserver := httpserver.NewRestfulServer(config.GlobalConfig.HttpAddr(), serve)
	apiserver.Init()
	go apiserver.Run()

	<-utils.WaitForSignal()

	cancel()
	apiserver.Shutdown()

}

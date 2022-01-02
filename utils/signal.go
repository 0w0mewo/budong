package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitForSignal() chan os.Signal {
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	return quit
}

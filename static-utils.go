package nlibgo

import (
	"os"
	"os/signal"
	"syscall"
)

func Wait() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

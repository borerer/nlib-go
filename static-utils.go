package nlibgo

import (
	"fmt"
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

func Safe(f func()) (err error) {
	defer func() {
		t := recover()
		if t != nil {
			err = fmt.Errorf("%+v", t)
		}
	}()
	f()
	return
}

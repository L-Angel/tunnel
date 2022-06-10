package main

import (
	"github.com/l-angel/tunnel/bootstrap"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	kill := make(chan bool, 1)
	b := bootstrap.Bootstrap{}
	b.Boot()
	hold(kill)
}

func hold(kill <-chan bool) {
	for {
		select {
		case <-kill:
			runtime.Goexit()
		}
	}
}

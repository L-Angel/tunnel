package main

import (
	"github.com/l-angel/tunnel/bootstrap"
	"time"
)

func main() {
	b := bootstrap.Bootstrap{}
	b.Boot()
	hold()
}

func hold() {
	for {
		time.Sleep(1 * time.Hour)
	}
}

package main

import (
	"github.com/m4t1t0/go-hotels-proxy/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}

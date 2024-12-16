package bootstrap

import (
	"github.com/m4t1t0/go-hotels-proxy/internal/platform/server"
)

const (
	port = 3000
)

func Run() error {
	srv := server.New(port)
	return srv.Run()
}

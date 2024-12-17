package bootstrap

import (
	"github.com/m4t1t0/go-hotels-proxy/internal/platform/server"
	"os"
	"strconv"
)

func Run() error {
	port, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
	srv := server.New(uint(port))
	return srv.Run()
}

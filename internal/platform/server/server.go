package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/m4t1t0/go-hotels-proxy/internal/platform/server/handler/home"
	"log"
)

type Server struct {
	port uint
	app  *fiber.App
}

func New(port uint) Server {
	srv := Server{
		app:  fiber.New(),
		port: port,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on port", s.port)
	return s.app.Listen(fmt.Sprintf(":%d", s.port))
}

func (s *Server) registerRoutes() {
	s.app.Get("/", home.Handler())
}

package server

import (
	"github.com/gofiber/fiber/v3"
)

const API_PREFIX = "/api/v1"

type ServerOptions struct{}

type APIServer struct {
}

func NewServer() (*APIServer, error) {
	server := &APIServer{}

	server.initServer()

	return server, nil
}

func (s *APIServer) initServer() error {
	app := s.createRouter()

	return app.Listen(":9090")
}

func (s *APIServer) createRouter() *fiber.App {
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	return app
}

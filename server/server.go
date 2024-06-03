package server

import (
	"strconv"

	"github.com/OhMinsSup/tavoli/config"
	"github.com/gofiber/fiber/v3"
)

const API_PREFIX = "/api/v1"

type APIServer struct {
	Config *config.Configuration
}

func NewServer(
	cfg *config.Configuration,
) (*APIServer, error) {
	server := &APIServer{
		Config: cfg,
	}

	server.initServer()

	return server, nil
}

func (s *APIServer) initServer() error {
	app := s.createRouter()

	port := strconv.Itoa(s.Config.Port)
	return app.Listen(":" + port)
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

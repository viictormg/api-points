package server

import (
	"api-points/config"
	"api-points/features"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	cfg     *config.Config
	Feature features.Feature
}

func NewServer(cfg *config.Config, Feature features.Feature) *Server {
	return &Server{
		cfg:     cfg,
		Feature: Feature,
	}
}

func (s Server) Run() error {

	app := fiber.New()

	s.RegisterRoute(app)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "HEAD,GET,POST,PUT,DELETE,PATCH",
		AllowCredentials: false,
	}))
	return app.Listen(s.cfg.Host)
}

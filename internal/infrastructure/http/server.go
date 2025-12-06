package http

import (
	"fmt"

	"go-clean-arch-template/internal/config"
	"go-clean-arch-template/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
	cfg *config.Config
}

func NewServer(cfg *config.Config, userUC *usecase.UserUsecase) *Server {
	app := fiber.New()

	RegisterRoutes(app, userUC)

	return &Server{app, cfg}
}

func (s *Server) Start() error {
	fmt.Println("Server started at", s.cfg.AppPort)
	return s.app.Listen(s.cfg.AppPort)
}

package http

import (
	"github.com/gofiber/fiber/v2"
	"go-clean-arch-template/internal/delivery/http"
	"go-clean-arch-template/internal/usecase"
)

func RegisterRoutes(app *fiber.App, userUC *usecase.UserUsecase) {
	userHandler := http.NewUserHandler(userUC)

	api := app.Group("/api")
	{
		api.Post("/users", userHandler.Create)
		api.Get("/users/:id", userHandler.GetByID)
	}
}
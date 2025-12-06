package http

import (
	"user-management/internal/delivery/http"
	"user-management/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, userUC *usecase.UserUsecase) {
	userHandler := http.NewUserHandler(userUC)

	api := app.Group("/api")
	{
		api.Post("/users", userHandler.Create)
		api.Get("/users/:id", userHandler.GetByID)
	}
}

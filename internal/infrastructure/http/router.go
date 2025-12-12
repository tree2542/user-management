package http

import (
	"user-management/internal/delivery/http"
	"user-management/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, 
	userUC *usecase.UserUsecase,
	loginUC *usecase.LoginUsecase,
	) {

	userHandler := http.NewUserHandler(userUC)
	loginHandler := http.NewLoginUsecase(loginUC)

	api := app.Group("/api")
	{
		api.Post("/users", userHandler.Create)
		api.Post("/login", loginHandler.Login)
		// api.Get("/users/:id", userHandler.GetByID)
	}
}

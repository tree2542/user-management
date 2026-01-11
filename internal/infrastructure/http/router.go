package http

import (
	"user-management/internal/delivery/http"
	"user-management/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RegisterRoutes(app *fiber.App,
	userUC *usecase.UserUsecase,
	loginUC *usecase.LoginUsecase,
) {

	userHandler := http.NewUserHandler(userUC)
	loginHandler := http.NewLoginUsecase(loginUC)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	api := app.Group("/api")
	{
		api.Post("/register", userHandler.Register)
		api.Post("/login", loginHandler.Login)
		// api.Get("/users/:id", userHandler.GetByID)
	}
}

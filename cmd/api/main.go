package main

import (
	"log"

	"user-management/internal/config"
	"user-management/internal/infrastructure/db"
	"user-management/internal/infrastructure/http"
	"user-management/internal/infrastructure/logger"
	"user-management/internal/repository"
	"user-management/internal/usecase"
	// "user-management/internal/usecase/port"
)

func main() {
	// fmt.Println("test")
	cfg := config.Load()

	logg := logger.New()

	postgres := db.NewPostgres(cfg)

	userRepo := repository.NewUserRepository(postgres)
	userUC := usecase.NewUserUsecase(userRepo, logg)

	server := http.NewServer(cfg, userUC)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

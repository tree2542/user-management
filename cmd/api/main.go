package main

import (
	"log"

	"go-clean-arch-template/internal/config"
	"go-clean-arch-template/internal/infrastructure/db"
	"go-clean-arch-template/internal/infrastructure/http"
	"go-clean-arch-template/internal/infrastructure/logger"
	"go-clean-arch-template/internal/repository"
	"go-clean-arch-template/internal/usecase"
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

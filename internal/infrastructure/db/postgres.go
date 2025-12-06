package db

import (
	"log"

	"go-clean-arch-template/internal/config"
	"go-clean-arch-template/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.PostgresDsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect postgres: ", err)
	}

	db.AutoMigrate(&domain.User{})

	return db
}
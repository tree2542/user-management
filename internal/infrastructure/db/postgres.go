package db

import (
	"log"

	"user-management/internal/config"
	"user-management/internal/domain"

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

package repository

import (
	"context"

	"user-management/internal/domain"

	"gorm.io/gorm"
)

// # Repository Interfaces
// db cmd
type UserRepository interface {
	Create(ctx context.Context, u *domain.User) error
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, u *domain.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

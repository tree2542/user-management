package usecase

import (
	"context"

	"user-management/internal/domain"
	"user-management/internal/infrastructure/logger"
	"user-management/internal/repository"
)

// # Application Business Rules
// handler
type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository, log *logger.Logger) *UserUsecase {
	return &UserUsecase{repo}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, u *domain.User) error {
	return uc.repo.Create(ctx, u)
}

func (uc *UserUsecase) GetUser(ctx context.Context, id uint) (*domain.User, error) {
	return uc.repo.GetByID(ctx, id)
}

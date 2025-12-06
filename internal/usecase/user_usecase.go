package usecase

import (
	"context"

	"go-clean-arch-template/internal/domain"
	"go-clean-arch-template/internal/repository"
	"go-clean-arch-template/internal/infrastructure/logger"
)

//# Application Business Rules
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
package usecase

import (
	"context"
	"fmt"
	"user-management/internal/delivery/dto"
	"user-management/internal/domain"
	"user-management/internal/repository"
	"user-management/internal/usecase/port"
)

type LoginUsecase struct {
	userRepo repository.UserRepository
	logger   port.Logger
}

func NewLoginUsecase(
	repo repository.UserRepository,
	log port.Logger,
) *LoginUsecase {
	return &LoginUsecase{repo, log}
}

func (uc *LoginUsecase) Login(ctx context.Context, req *dto.LoginRequest) (*domain.User, error) {

	getUser, err := uc.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		uc.logger.Error("user not found")
		return nil, fmt.Errorf("404")
	}

	return getUser, nil
}

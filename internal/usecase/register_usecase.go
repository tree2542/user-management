package usecase

import (
	"context"
	"fmt"
	"user-management/internal/delivery/dto"

	"user-management/internal/domain"
	"user-management/internal/repository"
	"user-management/internal/usecase/port"
)

// # Application Business Rules
// handler
type UserUsecase struct {
	userRepo repository.UserRepository
	logger   port.Logger
}

func NewUserUsecase(
	repo repository.UserRepository,
	log port.Logger,
) *UserUsecase {
	return &UserUsecase{repo, log}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, req *dto.UserCreateRequest) error {
	fmt.Println("This business logic")

	user := domain.User{
		Username: req.Username,
		FirstName: req.FirstName,
		LastName: req.LastName,
		Password: req.Password, //Hash in the future
		Email:    req.Email,
		Role: "",
	}
	user.SetModel()

	err := uc.userRepo.Create(ctx, &user)
	if err != nil {
		uc.logger.Error("failed to create user", map[string]interface{}{
			"username": user.Username,
			"error":    err.Error(),
		})
		return fmt.Errorf("usecase: create user: %w", err)
	}

	return nil
}

// func (uc *UserUsecase) GetUser(ctx context.Context, id uint) (*domain.User, error) {
// 	return uc.userRepo.GetByID(ctx, id)
// }

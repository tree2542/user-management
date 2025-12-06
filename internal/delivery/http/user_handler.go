package http

import (
	"strconv"

	"user-management/internal/delivery/dto"
	"user-management/internal/domain"
	"user-management/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

// Controller
type UserHandler struct {
	uc *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req dto.UserCreateRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user := &domain.User{
		Username: req.Username,
		Email:    req.Email,
	}

	if err := h.uc.CreateUser(c.Context(), user); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(fiber.Map{"message": "created"})
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user, err := h.uc.GetUser(c.Context(), uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.JSON(user)
}

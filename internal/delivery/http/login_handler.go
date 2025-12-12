package http

import (
	"fmt"
	"user-management/internal/delivery/dto"
	"user-management/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type LoginHandler struct {
	uc *usecase.LoginUsecase
}

func NewLoginUsecase(uc *usecase.LoginUsecase) *LoginHandler {
	return &LoginHandler{uc}
}

func (h *LoginHandler) Login(c *fiber.Ctx) error {
	fmt.Println("LOGINNNN")
	var req dto.LoginRequest
	var resp dto.ResponseModel

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	respData, err := h.uc.Login(c.Context(), &req)
	if err != nil {
		if err.Error() == "404" {
			resp.SetNotFoundError("ไม่พบข้อมูล username")
			return c.Status(fiber.StatusNotFound).JSON(resp)
		} else {
			resp.SetInternalServerError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(resp)
		}
	}

	setResponse := dto.LoginResponse{
		Username:  respData.Username,
		FirstName: respData.FirstName,
		LastName:  respData.LastName,
		Role:      respData.Role,
	}

	resp.SetSuccess(setResponse)
	return c.Status(fiber.StatusOK).JSON(resp)
}

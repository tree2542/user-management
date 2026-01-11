package http

import (
	"user-management/internal/delivery/dto"
	// "user-management/internal/domain"
	"user-management/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

// Controller (ถ้า Clean จะเรียกว่าเป็น Handler)
type UserHandler struct {
	uc *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req dto.UserCreateRequest
	var resp dto.ResponseModel

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := h.uc.CreateUser(c.Context(), &req); err != nil {
		//แยกเงื่อนไขการแบ่ง error ในนี้
		resp.SetInternalServerError("พบข้อผิดพลาด")
		return c.Status(500).JSON(resp)
	}

	setResponse := dto.ResultModel{
		ResultTest: "ลงทะเบียนสำเร็จ",
	}

	resp.SetSuccess(setResponse)
	return c.Status(200).JSON(resp)
}

// func (h *UserHandler) GetByID(c *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(c.Params("id"))

// 	user, err := h.uc.GetUser(c.Context(), uint(id))
// 	if err != nil {
// 		return c.Status(404).JSON(err.Error())
// 	}

// 	return c.JSON(user)
// }

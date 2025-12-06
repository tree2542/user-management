package dto

// req controller
type UserCreateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
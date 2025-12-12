package dto

// req controller
type UserCreateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ResultModel struct{
	ResultTest string `json:"result_test"`
}
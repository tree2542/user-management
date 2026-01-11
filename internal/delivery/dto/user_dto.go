package dto

// req controller
type UserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email    string `json:"email"`
}

type ResultModel struct{
	ResultTest string `json:"result_test"`
}
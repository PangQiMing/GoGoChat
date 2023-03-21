package dto

type RegisterUserDTO struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Sex      string `json:"sex"`
	Age      uint   `json:"age"`
}

type LoginUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

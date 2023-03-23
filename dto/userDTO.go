package dto

type RegisterUserDTO struct {
	GoGoID   uint64 `json:"go_go_id"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Sex      string `json:"sex"`
	Age      uint   `json:"age"`
}

type LoginUserDTO struct {
	GoGoID   uint64 `json:"go_go_id"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	AvatarURL string `json:"avatar_url"`
	Nickname  string `json:"nickname"`
	Sex       string `json:"sex"`
	Age       uint   `json:"age"`
}

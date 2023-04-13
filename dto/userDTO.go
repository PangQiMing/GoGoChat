package dto

type RegisterUserDTO struct {
	GoGoID   uint64 `json:"go_go_id" form:"go_go_id"`
	Nickname string `json:"nickname" form:"nickname"`
	Password string `json:"password" form:"password"`
	Sex      string `json:"sex" form:"sex"`
	Age      string `json:"age" form:"age"`
}

type LoginUserDTO struct {
	GoGoID   uint64 `json:"go_go_id"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	Nickname     string `json:"nickname"`
	Sex          string `json:"sex"`
	Age          string `json:"age"`
	introduction string `json:"introduction"`
}

type UpdateUserPasswdDTO struct {
	GoGoID      uint64 `json:"-"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

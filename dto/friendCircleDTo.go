package dto

type CircleCreateDTO struct {
	GoGoID   uint64 `json:"go_go_id"`
	Nickname string `json:"nickname"`
	Content  string `json:"content"`
	Picture  string `json:"picture"`
}

type DeleteFriendCircleDTO struct {
	FriendCircleID uint64 `json:"friend_circle"`
}

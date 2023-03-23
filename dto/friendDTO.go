package dto

type AddFriendDTO struct {
	GoGoID   uint64 `json:"go_go_id"`
	FriendID uint64 `json:"friend_id"`
}

type DeleteFriendDTO struct {
	GoGoID   uint64 `json:"go_go_id"`
	FriendID uint64 `json:"friend_id"`
}

type AcceptFriendDTO struct {
	GoGoID   uint64 `json:"go_go_id"`
	FriendID uint64 `json:"friend_id"`
}

type RejectFriendDTO struct {
	GoGoID   uint64 `json:"go_go_id"`
	FriendID uint64 `json:"friend_id"`
}

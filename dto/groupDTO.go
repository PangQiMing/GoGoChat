package dto

type CreateGroupDTO struct {
	Name         string `json:"name"`
	Announcement string `json:"announcement"`
}

type UpdateGroupDTO struct {
	GroupID      uint64 `json:"group_id"`
	Name         string `json:"name"`
	Announcement string `json:"announcement"`
}

type DeleteGroupDTO struct {
	GroupID uint64 `json:"group_id"`
}

type JoinGroupDTO struct {
	GroupID      uint64 `json:"group_id"`
	GroupOwnerID uint64 `json:"group_owner_id"`
	MemberID     uint64 `json:"-"`
}

type AcceptJoinGroupDTO struct {
	GroupID      uint64 `json:"group_id"`
	GroupOwnerID uint64 `json:"-"`
	MemberID     uint64 `json:"member_id"`
}

type RejectJoinGroupDTO struct {
	GroupID  uint64 `json:"group_id"`
	MemberID uint64 `json:"member_id"`
}

type SearchGroupDTO struct {
	GroupID string `json:"group_id"`
}

type DeleteGroupMemberDTO struct {
	GroupID string `json:"group_id"`
}

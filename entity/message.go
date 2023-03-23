package entity

type Message struct {
	Avatar   string `json:"avatar"`   //头像
	Nickname string `json:"nickname"` //昵称
	Content  string `json:"content"`  //内容
}

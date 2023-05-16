package entity

type Message struct {
	Avatar       string `gorm:"not null" json:"avatar"`
	FromUsername string `gorm:"varchar(50)" json:"from_username"`
	From         string `gorm:"uint;not null" json:"from"`
	To           string `gorm:"uint;not null" json:"to"`
	Content      string `gorm:"long text" json:"content"`
	MessageType  string `gorm:"not null" json:"message_type"`
	Time         string `gorm:"not null" json:"time"`
}

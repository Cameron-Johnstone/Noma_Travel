// models/message.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SenderID   uint
	Sender     User `gorm:"foreignKey:SenderID"`
	ReceiverID uint
	Receiver   User   `gorm:"foreignKey:ReceiverID"`
	Content    string `gorm:"not null"`
	SentAt     time.Time
}

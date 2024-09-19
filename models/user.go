// models/user.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username          string `gorm:"uniqueIndex;not null"`
	Email             string `gorm:"uniqueIndex;not null"`
	PasswordHash      string `gorm:"not null"`
	FullName          string
	Bio               string
	ProfilePictureURL string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

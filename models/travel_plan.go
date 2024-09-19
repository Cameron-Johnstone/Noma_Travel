// models/travel_plan.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type TravelPlan struct {
	gorm.Model
	UserID      uint
	User        User   `gorm:"foreignKey:UserID"`
	Title       string `gorm:"not null"`
	Description string
	StartDate   time.Time `gorm:"not null"`
	EndDate     time.Time `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

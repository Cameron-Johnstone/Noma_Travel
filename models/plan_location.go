// models/plan_location.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type PlanLocation struct {
	gorm.Model
	PlanID        uint
	TravelPlan    TravelPlan `gorm:"foreignKey:PlanID"`
	LocationID    uint
	Location      Location  `gorm:"foreignKey:LocationID"`
	ArrivalDate   time.Time `gorm:"not null"`
	DepartureDate time.Time `gorm:"not null"`
}

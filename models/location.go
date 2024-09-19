// models/location.go
package models

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Name      string  `gorm:"not null"`
	Latitude  float64 `gorm:"not null"`
	Longitude float64 `gorm:"not null"`
}

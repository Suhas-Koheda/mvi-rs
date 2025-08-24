package models

import "gorm.io/gorm"

type Showtime struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	MovieID uint `gorm:"not null;index"`
	HallID  uint `gorm:"not null;index"`
}

package models

import "gorm.io/gorm"

type Hall struct {
	gorm.Model
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"not null"`
	Capacity  int        `gorm:"not null"`
	Showtimes []Showtime `gorm:"foreignKey:HallID"`
	Seats     []Seat     `gorm:"foreignKey:HallID"`
}

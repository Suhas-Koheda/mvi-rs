package models

import (
	"time"

	"gorm.io/gorm"
)

type Showtime struct {
	gorm.Model
	ID           uint          `gorm:"primaryKey"`
	MovieID      uint          `gorm:"not null;index"`
	HallID       uint          `gorm:"not null;index"`
	StartTime    time.Time     `gorm:"not null;index"`
	EndTime      time.Time     `gorm:"not null;index"`
	TicketPrice  float64       `gorm:"type:decimal(10,2);not null"`
	Movie        Movie         `gorm:"foreignKey:MovieID"`
	Hall         Hall          `gorm:"foreignKey:HallID"`
	Reservations []Reservation `gorm:"foreignKey:ShowtimeID"`
	Seats        []Seat        `gorm:"foreignKey:ShowtimeID"`
}

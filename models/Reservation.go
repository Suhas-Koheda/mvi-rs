package models

import (
	"time"

	"gorm.io/gorm"
)

type ReservationStatus string

const (
	ReservationStatusActive    ReservationStatus = "active"
	ReservationStatusCancelled ReservationStatus = "cancelled"
	ReservationStatusPending   ReservationStatus = "pending"
)

type Reservation struct {
	gorm.Model
	ID            uint              `gorm:"primaryKey"`
	UserID        uint              `gorm:"not null;index"`
	ShowtimeID    uint              `gorm:"not null;index"`
	NumSeats      int               `gorm:"not null"`
	TotalPrice    float64           `gorm:"type:decimal(10,2);not null"`
	ReservedAt    time.Time         `gorm:"autoCreateTime"`
	Status        ReservationStatus `gorm:"type:varchar(20);default:'active'"`
	User          User              `gorm:"foreignKey:UserID"`
	Showtime      Showtime          `gorm:"foreignKey:ShowtimeID"`
	ReservedSeats []Seat            `gorm:"foreignKey:ReservationID"`
}

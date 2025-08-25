package models

import "gorm.io/gorm"

type Seat struct {
	gorm.Model
	ID            uint        `gorm:"primaryKey"`
	ShowtimeID    uint        `gorm:"not null;uniqueIndex:idx_showtime_seat"`
	HallID        uint        `gorm:"not null"`
	SeatNumber    string      `gorm:"not null;uniqueIndex:idx_showtime_seat"`
	IsReserved    bool        `gorm:"not null;default:false"`
	ReservationID *uint       `gorm:"index"`
	Showtime      Showtime    `gorm:"foreignKey:ShowtimeID"`
	Hall          Hall        `gorm:"foreignKey:HallID"`
	Reservation   Reservation `gorm:"foreignKey:ReservationID"`
}

package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	ID              uint   `gorm:"primaryKey"`
	Title           string `gorm:"not null"`
	Description     string
	PosterImageURL  string `gorm:"column:poster_image_url"`
	Genre           string
	DurationMinutes int        `gorm:"not null"`
	Showtimes       []Showtime `gorm:"foreignKey:MovieID"`
}

package models

import "gorm.io/gorm"

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	FirstName    string
	LastName     string
	Role         Role          `gorm:"type:varchar(10);default:'user'"`
	Reservations []Reservation `gorm:"foreignKey:UserID"`
}

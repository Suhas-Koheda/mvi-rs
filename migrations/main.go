package main

import (
	initializers "github.com/suhas-koheda/mvi-rs/initialisers"
	"github.com/suhas-koheda/mvi-rs/models"
)

func init() {
	initializers.InitialiseDatabase()
}

func main() {
	err := initializers.DB.AutoMigrate(
		&models.Reservation{}, &models.Showtime{}, &models.Seat{},
		&models.User{}, &models.Hall{}, &models.Movie{},
	)
	if err != nil {
		panic(err)
	}
}

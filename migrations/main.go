package main

import initializers "github.com/suhas-koheda/mvi-rs/initialisers"

func init() {
	initializers.InitialiseDatabase()
}

func main() {
	err := initializers.DB.AutoMigrate()
	if err != nil {
		panic(err)
	}
}

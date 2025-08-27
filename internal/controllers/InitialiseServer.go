package controllers

import (
	"gofr.dev/pkg/gofr"
)

var app *gofr.App

func initialiseServer() {
	app = gofr.New()
}

func registerControllers() {
	mainPageController()
	authControllers()
}

func StartServer() {
	initialiseServer()
	registerControllers()
	app.Run()
}

package controllers

import (
	"gofr.dev/pkg/gofr"
)

var app *gofr.App

func InitialiseServer(){ 
	app=gofr.New()
}

func Controllers(){
	    MainPageController() 
}

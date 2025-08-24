package controllers

import "gofr.dev/pkg/gofr"

func mainPageController()  {
	app.GET("/",mainPageHandler)
}

func mainPageHandler(cts *gofr.Context) (any,error){
	return "Hello World!",nil
}

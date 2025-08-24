package main

import (
	"fmt"

	"github.com/suhas-koheda/mvi-rs/controllers"
	"github.com/suhas-koheda/mvi-rs/env"
)

func main()  {
		env.LoadEnv()
	fmt.Println(env.EnvKey)
		controllers.InitialiseServer() 
		controllers.Controllers() 
}

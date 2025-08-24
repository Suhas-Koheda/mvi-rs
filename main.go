package main

import (
	"fmt"

	"github.com/suhas-koheda/mvi-rs/controllers"
	"github.com/suhas-koheda/mvi-rs/env"
	initializers "github.com/suhas-koheda/mvi-rs/initialisers"
)

func main()  {
		env.LoadEnv()
	fmt.Println(env.EnvKey)
		controllers.StartServer() 
		initializers.InitialiseDatabase() 
}

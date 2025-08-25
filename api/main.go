package main

import (
	"fmt"

	"github.com/suhas-koheda/mvi-rs/internal/controllers"
	"github.com/suhas-koheda/mvi-rs/internal/oauth"
	"github.com/suhas-koheda/mvi-rs/pkg/env"
	initializers "github.com/suhas-koheda/mvi-rs/pkg/initialisers"
)

func main() {
	loadenv.LoadEnv()
	fmt.Println(loadenv.EnvKey)
	controllers.StartServer()
	initializers.InitialiseDatabase()
	oauth.InitialiseOauth()
}

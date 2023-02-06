package main

import (
	"fmt"
	"log"

	"github.com/redbeestudios/go-seed/cmd"
	"github.com/redbeestudios/go-seed/pkg"
)

func main() {
	env, err := pkg.NewEnv("dev")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.LUTC)

	if err != nil {
		panic(fmt.Sprintf("error creating environment: %s", err.Error()))
	}

	config := cmd.InitConfig(env)
	deps := cmd.InitDependencies(config)
	router := cmd.InitRoutes(deps)

	cmd.StartServer(config, router)
}

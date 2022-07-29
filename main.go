package main

import (
	"github.com/labstack/echo/v4"
	"global.p2p.api/agents"
	"global.p2p.api/gp2p"
	"log"
)

func main() {
	env, err := gp2p.NewEnv(".env")
	if err != nil {
		log.Fatal(err)
	}

	app := gp2p.New(env, echo.New())

	app.EnableModule(agents.New())

	app.StartServer()
}

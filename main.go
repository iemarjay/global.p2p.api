package main

import (
	"github.com/labstack/echo/v4"
	"global.p2p.api/agents"
	"global.p2p.api/app"
	"log"
)

func main() {
	env, err := app.NewEnv(".env")
	if err != nil {
		log.Fatal(err)
	}

	app := app.New(env, echo.New())

	app.EnableModule(agents.New())

	app.StartServer()
}

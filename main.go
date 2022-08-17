package main

import (
	"github.com/labstack/echo/v4"
	"global.p2p.api/agents"
	"global.p2p.api/app"
	"global.p2p.api/invoices"
	"log"
)

func main() {
	env, err := app.NewEnv(".env")
	if err != nil {
		log.Fatal(err)
	}

	a := app.New(env, echo.New())

	a.EnableModule(agents.New())
	a.EnableModule(invoices.New())

	a.StartServer()
}

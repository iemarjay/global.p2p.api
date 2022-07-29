package gp2p

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	database2 "global.p2p.api/gp2p/database"
	"global.p2p.api/gp2p/http"
)

type Gp2p struct {
	echo    *echo.Echo
	env     *Env
}

type Module interface {
	Init(app Gp2p)
}

func (app Gp2p) EnableModule(module Module) Gp2p {
	module.Init(app)

	return app
}

func (app Gp2p) Env() *Env {
	return app.env
}

func (app Gp2p) Router() *echo.Echo {
	return app.echo
}

func (app Gp2p) Database() database2.Database {
	env := app.env
	return database2.NewMongo(env.Get("DATABASE_URL"), env.Get("DATABASE_NAME"))
}

func (app Gp2p) StartServer() {
	address := ":" + app.env.GetOrDefault("APP_PORT", "8001")

	//app.echo.GET("/", func (c echo.Context) error {
	//	return c.JSON(200, app.echo.Routes())
	//})

	app.echo.Logger.Fatal(app.echo.Start(address))
}

func New(env *Env, e *echo.Echo) *Gp2p {
	e.Validator = http.NewAppValidator(validator.New())

	return &Gp2p{
		env: env,
		echo: e,
	}
}

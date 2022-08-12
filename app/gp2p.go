package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	database2 "global.p2p.api/app/database"
	"global.p2p.api/app/fileStorage"
	"global.p2p.api/app/http"
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

func (app Gp2p) Echo() *echo.Echo {
	return app.echo
}

func (app Gp2p) Database() database2.Database {
	env := app.env
	return database2.NewMongo(env.Get(DATABASE_URL), env.Get(DATABASE_NAME))
}

func (app Gp2p) FileSystem() *fileStorage.FileStorage {
	baseUrl := app.env.Get(APP_BASE_URL)
	rootDir := app.env.GetOrDefault(PUBLIC_ROOT_DIR, PUBLIC_ROOT_DIR_VALUE)
	opts := fileStorage.NewLocalDiskOpts(rootDir, baseUrl, app.env.GetOrDefault(PUBLIC_PATH_PREFIX, PUBLIC_PATH_PREFIX_VALUE))
	d := fileStorage.NewPublicDisk(opts)

	return fileStorage.NewFileSystem(d)
}

func (app Gp2p) StartServer() {
	address := ":" + app.env.GetOrDefault(APP_PORT, "8001")
	app.echo.Static(app.env.GetOrDefault(PUBLIC_PATH_PREFIX, PUBLIC_PATH_PREFIX_VALUE),
		app.env.GetOrDefault(PUBLIC_ROOT_DIR, PUBLIC_ROOT_DIR_VALUE))

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

package adverts

import (
	"github.com/labstack/echo/v4"
	"global.p2p.api/adverts/controller"
)

func New(echo *echo.Echo) {
	advertEcho := echo.Group("advert")

	advertEcho.POST("create", controller.CreateAdvert)
}

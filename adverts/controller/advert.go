package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateAdvert(c echo.Context) error {
	return c.JSON(http.StatusCreated, actions)
}

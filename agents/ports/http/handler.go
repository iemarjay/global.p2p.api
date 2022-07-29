package http

import (
	"github.com/labstack/echo/v4"
	"global.p2p.api/agents/services"
	"net/http"
)

type Handler struct {
	service *services.AgentService
}

func New(as *services.AgentService) *Handler {
	return &Handler{
		service: as,
	}
}

func (h Handler) Register(e *echo.Echo) {
	router := e.Group("/agent")
	
	router.POST("/register", h.register())
}

func (h Handler) register() echo.HandlerFunc {
	return func(c echo.Context) error {
		fields := &AgentRegisterRequest{}

		if err := c.Bind(fields); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if err := c.Validate(fields); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		asd := fields.toAgentServiceData()
		agent, err := h.service.RegisterAgent(asd)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, agent)
	}
}
package http

import (
	"github.com/labstack/echo/v4"
	"global.p2p.api/agents/dtos"
	"global.p2p.api/agents/services"
	"global.p2p.api/app"
	appHttp "global.p2p.api/app/http"
	"net/http"
)

type Handler struct {
	service             *services.AgentService
	loginService        *services.LoginService
	verificationService  *services.AgentVerificationService
}

func New(as *services.AgentService, loginService *services.LoginService,
	vs *services.AgentVerificationService) *Handler {
	return &Handler{
		service:             as,
		loginService:        loginService,
		verificationService:  vs,
	}
}

func (h Handler) RegisterRoutes(a app.Gp2p) {
	router := a.Echo().Group("/agent")

	router.POST("/register", h.register())
	router.POST("/login", h.login())

	router.POST("/verification/send-code", h.sendVerificationCode(), appHttp.AuthMiddleware())
	router.POST("/verification", h.verifyToken(), appHttp.AuthMiddleware())

	router.POST("/kyc", h.verifyToken(), appHttp.AuthMiddleware())

}

func (h Handler) register() echo.HandlerFunc {
	return func(c echo.Context) error {
		fields := &AgentRegisterRequest{}

		if err := c.Bind(fields); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if err := c.Validate(fields); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		asd := fields.toAgentServiceData()
		agent, err := h.service.RegisterAgent(asd)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, agent)
	}
}

func (h *Handler) login() echo.HandlerFunc {
	return func(c echo.Context) error {
		fields := &AgentLoginRequest{}

		if err := c.Bind(fields); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if err := c.Validate(fields); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		data := fields.toAgentServiceData()
		login, err := h.loginService.Login(data)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, login)
	}
}

func (h Handler) sendVerificationCode() echo.HandlerFunc {
	return func(c echo.Context) error {
		claim := appHttp.ContextToClaim(c)
		input := &dtos.AgentVerificationDTO{ID: claim.Id}
		c.Bind(&input)

		err := h.verificationService.SendAgentVerificationCode(input)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, struct {
			Status string `json:"status"`
		}{Status: "success"})
	}
}

func (h Handler) verifyToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		claim := appHttp.ContextToClaim(c)
		input := &dtos.AgentVerificationDTO{ID: claim.Id}
		c.Bind(&input)

		err, ve := h.verificationService.VerifyAgent(input)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}

		if ve != nil {
			return echo.NewHTTPError(http.StatusNotFound, ve)
		}

		return c.JSON(http.StatusOK, struct {
			Status string `json:"status"`
		}{Status: "success"})
	}
}

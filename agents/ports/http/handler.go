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
	service         *services.Agent
	loginService    *services.Login
	verification    *services.AgentVerification
	kyc             *services.AgentKyc
}

func New(as *services.Agent, l *services.Login, vs *services.AgentVerification,
	kyc *services.AgentKyc) *Handler {
	return &Handler{
		service:      as,
		loginService: l,
		verification: vs,
		kyc:          kyc,
	}
}

func (h Handler) RegisterRoutes(a app.Gp2p) {
	router := a.Echo().Group("/agent")

	router.POST("/register", h.register())
	router.POST("/login", h.agentLogin())

	router.POST("/verification/send-code", h.sendVerificationCode(), appHttp.AuthMiddleware())
	router.POST("/verification", h.verifyToken(), appHttp.AuthMiddleware())

	router.POST("/kyc", h.agentKyc(), appHttp.AuthMiddleware())

}

func (h Handler) register() echo.HandlerFunc {
	return func(c echo.Context) error {
		fields := &dtos.AgentRegisterRequest{}
		if err := c.Bind(fields); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if err := c.Validate(fields); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		asd := fields.ToAgentServiceData()
		agent, err := h.service.RegisterAgent(asd)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, agent)
	}
}

func (h *Handler) agentLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		fields := &dtos.AgentLoginRequest{}

		if err := c.Bind(fields); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if err := c.Validate(fields); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		data := fields.ToAgentServiceData()
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

		err := h.verification.SendAgentVerificationCode(input)
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
		err := c.Bind(&input)
		if err != nil {
			return err
		}

		err, ve := h.verification.VerifyAgent(input)
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

func (h Handler) agentKyc() echo.HandlerFunc {
	return func(c echo.Context) error {
		fields := &dtos.KycRequest{}
		err := h.inputBind(c, fields)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if err := c.Validate(fields); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		claim := appHttp.ContextToClaim(c)
		agent, err := h.kyc.UpdateKyc(claim.Id, fields.ToKycDto())
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, agent)
	}
}

func (h Handler) inputBind(c echo.Context, fields *dtos.KycRequest) error {
	if err := c.Bind(fields); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var err error
	fields.IdCard, err = c.FormFile("id_card")
	if http.ErrMissingFile == err {
		return err
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	fields.AddressProof, err = c.FormFile("address_proof")
	if http.ErrMissingFile == err {
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return nil
}

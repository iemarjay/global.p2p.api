package http

import (
	"github.com/labstack/echo/v4"
	"global.p2p.api/app"
	"global.p2p.api/invoices/services"
	"net/http"
	"syreclabs.com/go/faker"
	"time"
)

type Handler struct {
	service *services.Invoice
}

func NewHandler(service *services.Invoice) *Handler {
	return &Handler{service: service}
}

func (h Handler) Register(app app.Gp2p) {
	router := app.Echo().Group("invoice")

	router.GET("/:id", h.viewInvoice())
	router.GET("/new/fake", h.generateFakeInvoice())
}

func (h Handler) viewInvoice() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		invoice, err := h.service.ShowInvoice(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, invoice)
	}
}

func (h Handler) generateFakeInvoice() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := h.agentFakeInput()
		invoice, err := h.service.CreateInvoice(input)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, invoice)
	}
}

func (h Handler) agentFakeInput() *services.InvoiceDto {
	qty := uint(faker.Number().NumberInt(2))
	price := uint(faker.Number().NumberInt(3))
	amount := qty * price
	items := []services.InvoiceItem{
		{
			Description: faker.Lorem().Sentence(5),
			Quantity:    qty,
			UnitPrice:   price,
			Amount:      amount,
		},
	}

	return &services.InvoiceDto{
		No: uint(faker.Number().NumberInt(5)),
		From: services.InvoiceContact{
			Name:  faker.Name().Name(),
			Email: faker.Internet().Email(),
		},
		To: services.InvoiceContact{
			Name:  faker.Name().Name(),
			Email: faker.Internet().Email(),
		},
		Fiat:     "USD",
		Crypto:   "cUSD",
		Items:    items,
		Total:    amount,
		Subtotal: amount,
		Note:     faker.Lorem().Sentence(30),
		IssuedAt: faker.Time().Backward(time.Hour * 24 * 7 * 3),
		DueAt:    faker.Time().Forward(time.Hour * 24 * 7 * 3),
	}
}

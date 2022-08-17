package invoices

import (
	"global.p2p.api/app"
	"global.p2p.api/invoices/port/http"
	"global.p2p.api/invoices/repositries"
	"global.p2p.api/invoices/services"
)

type invoice struct {
	app app.Gp2p
}

func New() *invoice {
	return &invoice{}
}

func (i invoice) Init(app app.Gp2p) {
	i.app = app

	service := i.makeInvoiceService()

	handler := http.NewHandler(service)
	handler.Register(app)
}

func (i invoice) makeInvoiceService() *services.Invoice {
	r := i.makeRepository()
	return services.NewInvoice(r)
}

func (i invoice) makeRepository() *repositries.Invoice {
	db := i.app.Database()
	db.Table("invoice")
	return repositries.NewInvoice(db)
}
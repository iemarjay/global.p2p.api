package services

import "time"

type InvoiceRepo interface {
	GetInvoiceById(id string) (*InvoiceDto, error)
	CreateInvoice(data *InvoiceDto) (*InvoiceDto, error)
}

type InvoiceContact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type InvoiceItem struct {
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	UnitPrice   uint   `json:"unit_price"`
	Amount      uint   `json:"amount"`
	Tax         uint   `json:"tax"`
	Discount    string `json:"discount"`
}

type InvoiceDto struct {
	ID           string         `json:"id"`
	No           uint           `json:"no"`
	From         InvoiceContact `json:"from"`
	To           InvoiceContact `json:"to"`
	Fiat         string         `json:"fiat"`
	Crypto       string         `json:"currency"`
	IssuedAt     time.Time      `json:"issued_at"`
	DueAt        time.Time      `json:"due_at"`
	Items        []InvoiceItem  `json:"items"`
	Subtotal     uint           `json:"subtotal"`
	Total        uint           `json:"total"`
	CryptoAmount uint           `json:"crypto_amount"`
	Note         string         `json:"note"`
}

type Invoice struct {
	repository InvoiceRepo
}

func NewInvoice(r InvoiceRepo) *Invoice {
	return &Invoice{repository: r}
}

func (i Invoice) ShowInvoice(id string) (*InvoiceDto, error) {
	invoice, err := i.repository.GetInvoiceById(id)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (i Invoice) CreateInvoice(input *InvoiceDto) (*InvoiceDto, error) {
	invoice, err := i.repository.CreateInvoice(input)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

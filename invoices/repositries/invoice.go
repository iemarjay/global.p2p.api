package repositries

import (
	"global.p2p.api/app/database"
	"global.p2p.api/invoices/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type contact struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

type item struct {
	Description string `bson:"description"`
	Quantity    uint   `bson:"quantity"`
	UnitPrice   uint   `bson:"unit_price"`
	Amount      uint   `bson:"amount"`
	Tax         uint   `bson:"tax"`
	Discount    string `bson:"discount"`
}

type mongoInvoice struct {
	ID           string             `bson:"_id,omitempty"`
	No           uint               `bson:"no"`
	From         contact            `bson:"from"`
	To           contact            `bson:"to"`
	Fiat         string             `bson:"fiat"`
	Crypto       string             `bson:"currency"`
	IssuedAt     primitive.DateTime `bson:"issued_at"`
	DueAt        primitive.DateTime `bson:"due_at"`
	Items        []item             `bson:"items"`
	Subtotal     uint               `bson:"subtotal"`
	Total        uint               `bson:"total"`
	CryptoAmount uint               `bson:"crypto_amount"`
	Note         string             `bson:"note"`
}

func NewMongoInvoiceFromInvoiceDto(i *services.InvoiceDto) *mongoInvoice {
	var items []item
	for _, invoiceItem := range i.Items {
		items = append(items, item{
			Description: invoiceItem.Description,
			Quantity:    invoiceItem.Quantity,
			UnitPrice:   invoiceItem.UnitPrice,
			Amount:      invoiceItem.Amount,
			Tax:         invoiceItem.Tax,
			Discount:    invoiceItem.Discount,
		})
	}

	return &mongoInvoice{
		ID:           i.ID,
		No:           i.No,
		From:         contact{Name: i.From.Name, Email: i.From.Email},
		To:           contact{Name: i.To.Name, Email: i.To.Email},
		Fiat:         i.Fiat,
		Crypto:       i.Crypto,
		IssuedAt:     primitive.NewDateTimeFromTime(i.IssuedAt),
		DueAt:        primitive.NewDateTimeFromTime(i.DueAt),
		Items:        items,
		Subtotal:     i.Subtotal,
		Total:        i.Total,
		CryptoAmount: i.CryptoAmount,
		Note:         i.Note,
	}
}

func (i *mongoInvoice) toInvoiceDto() *services.InvoiceDto {
	var items []services.InvoiceItem
	for _, invoiceItem := range i.Items {
		items = append(items, services.InvoiceItem{
			Description: invoiceItem.Description,
			Quantity:    invoiceItem.Quantity,
			UnitPrice:   invoiceItem.UnitPrice,
			Amount:      invoiceItem.Amount,
			Tax:         invoiceItem.Tax,
			Discount:    invoiceItem.Discount,
		})
	}

	return &services.InvoiceDto{
		ID:           i.ID,
		No:           i.No,
		From:         services.InvoiceContact{Name: i.From.Name, Email: i.From.Email},
		To:           services.InvoiceContact{Name: i.To.Name, Email: i.To.Email},
		Fiat:         i.Fiat,
		Crypto:       i.Crypto,
		IssuedAt:     i.IssuedAt.Time(),
		DueAt:        i.DueAt.Time(),
		Items:        items,
		Subtotal:     i.Subtotal,
		Total:        i.Total,
		CryptoAmount: i.CryptoAmount,
		Note:         i.Note,
	}
}

type Invoice struct {
	db database.Database
}

func NewInvoice(db database.Database) *Invoice {
	return &Invoice{db: db}
}

func (i Invoice) GetInvoiceById(id string) (*services.InvoiceDto, error) {
	invoice := &mongoInvoice{}
	err := i.db.FindOneByID(id).Decode(invoice)
	if err != nil {
		return nil, err
	}

	return invoice.toInvoiceDto(), nil
}

func (i Invoice) CreateInvoice(data *services.InvoiceDto) (*services.InvoiceDto, error) {
	invoice := &mongoInvoice{}
	decoder, err := i.db.Insert(data)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(&invoice)
	if err != nil {
		return nil, err
	}

	return invoice.toInvoiceDto(), nil
}

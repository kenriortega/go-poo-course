package invoice

import (
	"github.com/kenriortega/go-poo-course/modulo3/pkg/customer"
	"github.com/kenriortega/go-poo-course/modulo3/pkg/invoiceitem"
)

// Invoice is the struct of invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   invoiceitem.Items
}

// New return a new Invoice
func New(country, city string, client customer.Customer, items invoiceitem.Items) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}

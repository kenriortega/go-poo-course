package main

import (
	"fmt"

	"github.com/kenriortega/go-poo-course/modulo3/pkg/customer"
	"github.com/kenriortega/go-poo-course/modulo3/pkg/invoice"
	"github.com/kenriortega/go-poo-course/modulo3/pkg/invoiceitem"
)

func main() {

	i := invoice.New(
		"Cu",
		"Hav",
		customer.New("Ka", "Calle 123", "5699898"),
		invoiceitem.NewItems(
			invoiceitem.New(1, "laptop", 12.34),
			invoiceitem.New(2, "keyboard", 212.34),
			invoiceitem.New(3, "audi", 56.34),
		),
	)

	fmt.Printf("%+v\n", i)
}

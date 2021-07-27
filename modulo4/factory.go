package main

import "fmt"

type PayMethod interface {
	Pay()
}

type Paypal struct{}

func (p *Paypal) Pay() {
	fmt.Println("Paypal")
}

type Cash struct{}

func (c *Cash) Pay() {
	fmt.Println("Cash")
}

type CreditCard struct{}

func (cd *CreditCard) Pay() {
	fmt.Println("CreditCard")
}

func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return &Paypal{}
	case 2:
		return &Cash{}
	case 3:
		return &CreditCard{}
	default:
		return nil
	}
}
func main3() {
	var method uint
	fmt.Println("Digite una opcion de pago")
	fmt.Println("\t 1: Paypal 2: Cash 3: CreditCard")

	_, err := fmt.Scanln(&method)
	if err != nil {
		panic("Ingrese un methodo valido")
	}

	if method > 3 {
		panic("debe digitar un method valido")
	}
	payMethod := Factory(method)
	payMethod.Pay()
}

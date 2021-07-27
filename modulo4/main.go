package main

import "fmt"

type Greeter interface {
	Greet()
}

type Byer interface {
	Bye()
}

type GreeterByer interface {
	Greeter
	Byer
}

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Println("Hola " + p.Name)
}
func (p Person) Bye() {
	fmt.Println("Adios " + p.Name)
}

func (p Person) String() string {
	return "Soy un str " + p.Name
}

type Text string

func (t Text) Greet() {
	fmt.Println("Hola soy " + t)
}
func (t Text) Bye() {
	fmt.Println("Adios soy " + t)
}

// func GreetAll(gs ...Greeter) {
// 	for _, g := range gs {
// 		g.Greet()
// 		fmt.Printf("\t Valor %v, Tipo: %T\n", g, g)
// 	}
// }

// func ByeAll(bs ...Byer) {
// 	for _, b := range bs {
// 		b.Bye()
// 		fmt.Printf("\t Valor %v, Tipo: %T\n", b, b)
// 	}
// }

func All(gb ...GreeterByer) {
	for _, b := range gb {
		b.Bye()
		b.Greet()
		fmt.Printf("\t Valor %v, Tipo: %T\n", b, b)
	}
}

func main() {

	p := Person{Name: "as"}
	var t Text = "Ks"
	// GreetAll(p, t)
	// ByeAll(p, t)
	All(p, t)
	fmt.Println(p)
}

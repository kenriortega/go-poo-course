package main

import "fmt"

type Storager interface {
	Get() string
	Set(string)
}

type Person1 struct {
	name string
}

func NewPerson(name string) *Person1 {
	return &Person1{name: name}
}

func (p *Person1) Get() string {
	return p.name
}

func (p *Person1) Set(name string) {
	p.name = name
}

func Exec(s Storager, name string) {
	s.Set(name)
	fmt.Println(s.Get())

}

func main2() {

	p := NewPerson("ka")
	Exec(p, "ka")
}

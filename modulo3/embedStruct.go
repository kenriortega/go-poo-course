package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

func NewPerson(name string, age uint) Person {
	return Person{Name: name, Age: age}
}

func (p Person) Greet() {
	fmt.Println("Hello")
}

type Employee struct {
	Person // embed data
	Salary float64
}

func NewEmployee(name string, age uint, salary float64) Employee {
	return Employee{NewPerson(name, age), salary}
}

func (e Employee) Payroll() {
	fmt.Println(e.Salary * 0.9)
}
func (e Employee) Greet() {
	fmt.Println("Hello from employee")
}

func main3() {
	e := NewEmployee("KA", 30, 2344)
	e.Greet()
	e.Payroll()
}

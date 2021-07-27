package main

import "fmt"

type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

// declaracion de alias
type myAlias = course

// definiciones de type
type newCourse course

type newBool bool

func (b newBool) String() string {
	if b {
		return "Verdadero"
	}
	return "Falso"
}
func main2() {
	c := newCourse{name: "Go"}
	fmt.Printf("El tipo es %T\n", c)

	var b newBool = true
	fmt.Println(b.String())
}

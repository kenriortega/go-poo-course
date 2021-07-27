package main

import (
	"fmt"
)

type exmpler interface {
	X()
}

func wrapper(i interface{}) {

	// Assertions
	// v, ok := i.(string)
	// if ok {
	// 	result := strings.ToUpper(v)
	// 	fmt.Printf("Valor: %v, Tipo: %T\n", result, result)
	// }

	// Type Switch
	switch v := i.(type) {
	case string:
		fmt.Printf("Valor: %v, Tipo: %T\n", v, v)
	}
}

func main4() {
	wrapper("as")
}

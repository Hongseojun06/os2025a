package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var name string
	// name = "Kim Inha"

	// var name = "Kim Inha"

	// name := "Kim Inha"
	// name := 2.71

	// name = "Kim Inha"
	// name = 2.71

	var name float64
	name = 2.71
	fmt.Println(name, reflect.TypeOf(name))
}

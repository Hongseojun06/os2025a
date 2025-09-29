package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string
	//name = "Kim Inha"
	//var name="Kim Inha"
	//name := 2.71
	name = "Kim inha"
	fmt.Println(name, reflect.TypeOf(name))
}

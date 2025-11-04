package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//rand.Seed(10)
	dice := rand.Intn(6) + 1
	fmt.Println(dice)

}

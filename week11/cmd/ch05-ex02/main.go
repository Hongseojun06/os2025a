package main

import (
	"fmt"
	"log"

	//"week10/pkg/keyboard"
	"github.com/headfirstgo/datafile" // go get github.com/headfirstgo/keyboard
)

func main() {

	weights, err := datafile.GetFloats("meatWeight.txt")
	if err != nil {
		log.Fatal(err)
	}

	hap := 0.0

	for i := 0; i < len(weights); i++ {
		hap = hap + weights[i]
	}

	weeks := float64(len(weights))
	fmt.Println("평균:", hap/weeks)
}

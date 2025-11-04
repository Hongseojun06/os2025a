package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Print("화씨 온도 입력:")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("화씨온도 %.2f도 섭씨온도는 %.2f도입니다.", fahrenheit, celsius)
}

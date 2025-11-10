package main

import "fmt"

func main() {
	subjects := []string{"Go", "Javascript", "python", "Linux"} //initialized by slice litered
	subjectsSlice := subjects[1:3]
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("==========")
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}

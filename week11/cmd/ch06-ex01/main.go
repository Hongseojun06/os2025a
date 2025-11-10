package main

import "fmt"

func main() {
	subjects := []string{"Go", "", "python"} //initialized by slice litered

	for _, subject := range subjects {
		fmt.Println(subject)
	}
}

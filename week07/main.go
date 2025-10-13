package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//shadowing
	//var fmt string = "inha"
	//var int int = 7
	//var k int =11
	//fmt.Println(int)

	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n') //ignore error
	//fmt.Println(err)
	if err != nil {
		log.Fatal(err) //report the erro and exit program
	}
	i = strings.TrimSpace(i)
	score, err := strconv.ParseFloat(i, 64)
	if err != nil {
		log.Fatal(err)
	}
	if score >= 60 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")

	}
	fmt.Println(score)

}

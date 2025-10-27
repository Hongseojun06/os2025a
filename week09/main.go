// Package keyboard reads user input from the keyboard.
package keyboard

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// GetFloat reads a floating-point number from the keyboard.
// It returns the number read and any error encountered.
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

}
func main() {
	fmt.Print("점수 입력:")
	score, err := GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	status := ""
	if score >= 60 {
		status = "합격"
	} else {
		status = "불합격"
	}
	fmt.Printf("%.2f점은 %s\n", score, status)
}

package main

import (
	"fmt"
)

func main() {
	var value int = 2

	if value == 1 {
		fmt.Println("Condition is true")
	} else if value == 2 {
		fmt.Println("Condition is true")
	} else {
		fmt.Println("Condition is false")
	}
}

// go run app.go

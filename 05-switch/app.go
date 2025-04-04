package main

import (
	"fmt"
)

func main() {
	value := 10

	switch value {
	case 1:
		fmt.Println("Condition is true")
	case 2:
		fmt.Println("Condition is true")
	case 3:
		fmt.Println("Condition is true")
	default:
		fmt.Println("Condition is false")
	}
}

// go run app.go

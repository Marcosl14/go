package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Value is", i)
	}

	fmt.Println("------------------------")
	fmt.Println("------------------------")

	fmt.Println("For as while")
	var j = 0
	for {
		fmt.Println("Value is", j)
		j++
		if j == 3 {
			fmt.Println("Continue")
			continue
		}

		if j == 4 {
			fmt.Println("Break")
			break
		}
	}

	fmt.Println("------------------------")
	fmt.Println("------------------------")

	fmt.Println("For as while do")

	var condition = 1

	for condition < 5 {
		fmt.Println("Value is", condition)
		condition++
	}
}

// go run app.go

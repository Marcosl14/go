package main

import "fmt"

func main() {
	switchStatement(10)
	switchStatement("Hello")
	switchStatement(true)

	intergerValidation(10)
	intergerValidation("Hello")
}

func switchStatement(value any) {
	switch value.(type) {
	case int:
		fmt.Println("Value is int")
	case string:
		fmt.Println("Value is string")
	case bool:
		fmt.Println("Value is bool")
	case float64:
		fmt.Println("Value is float64")
	}
}

func intergerValidation(value any) {
	integerValue, isInteger := value.(int)

	if isInteger {
		fmt.Println("Value " + fmt.Sprintf("%d", integerValue) + " is int")
	} else {
		fmt.Println("Value " + fmt.Sprintf("%d", integerValue) + " is not int")
	}
}

// go run app.go

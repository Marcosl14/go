package main

import "fmt"

type str string

func (s str) Log() {
	fmt.Println("Logging:", s)
}

func main() {
	var name str = "John Doe" // Custom type
	name.Log()                // Custom type method
}

// go run app.go

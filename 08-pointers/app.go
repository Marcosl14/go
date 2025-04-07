package main

import (
	"fmt"
)

func main() {
	// pointers are variables that store the memory address of another variable

	fmt.Println("POINTERS AS VALUES")
	var value int = 2
	var pointer *int = &value // pointer is a variable that stores the memory address of value (reference)
	fmt.Println("Value:", value)
	fmt.Println("Pointer:", pointer)

	fmt.Println("------------------------")
	fmt.Println("------------------------")

	fmt.Println("REGULAR FUNCTIONS -> creates copies of arguments (see different memory addresses -> pointers)")
	doSomething(value)

	fmt.Println("------------------------")
	fmt.Println("------------------------")

	fmt.Println("POINTERS AS ARGUMENTS -> pass the memory address of the variable")
	doSomethingWithPointer(&value) // pass the meory address of the value to the function
}

// functions that takes values as arguments create a new scope (value int is a copy of the original value)
func doSomething(value int) {
	var pointer *int = &value
	fmt.Println("Function Value:", value)     // value is a copy of the original value
	fmt.Println("Function Pointer:", pointer) // pointer is the memory address of the copy
}

// functions that takes pointers as arguments do not create a new scope (pointer int is a reference to the original value)
// this means that if we change the value of the pointer, it will change the original value
func doSomethingWithPointer(pointer *int) {
	fmt.Println("Function Value:", *pointer)  // dereference the pointer to get the value
	fmt.Println("Function Pointer:", pointer) // pointer is the memory address of the original value
}

// go run app.go

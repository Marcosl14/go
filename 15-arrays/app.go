package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name      string
	Age       int
	CreatedAt time.Time
}

func main() {
	var arr1 = []int{1, 2, 3, 4, 5}
	var arr2 = []string{"a", "b", "c", "d", "e"}
	var arr3 = []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	var arr4 = []bool{true, false, true, false, true}
	var arr5 = []interface{}{1, "a", 2.2, true, false}
	var arr6 = []any{1, "a", 2.2, true, false}
	var arr7 = []Person{
		{Name: "John", Age: 30, CreatedAt: time.Now()},
		{Name: "Jane", Age: 25, CreatedAt: time.Now()},
		{Name: "Doe", Age: 35, CreatedAt: time.Now()},
	}

	fmt.Println("Array of int:", arr1)
	fmt.Println("Array of string:", arr2)
	fmt.Println("Array of float64:", arr3)
	fmt.Println("Array of bool:", arr4)
	fmt.Println("Array of interface{}:", arr5)
	fmt.Println("Array of any:", arr6)
	fmt.Println("Array of Person:", arr7)
}

// go run app.go

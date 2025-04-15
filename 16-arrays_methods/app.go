package main

import "fmt"

func main() {
	var arr1 = []int{1, 2, 3, 4, 5}            // Dynamic size array (unlimited size -> known as slice)
	var arr2 = [5]int{100, 200, 300, 400, 500} // Limited to only 5 elements -> known as array

	// The difference between a slice and an array is that a slice is a reference to an array, while an array is a value type
	// This means that when you pass a slice to a function, you are passing a reference to the original array, while when you pass an array to a function, you are passing a copy of the array

	fmt.Println(arr1)
	fmt.Println(arr2)

	// index
	fmt.Println(arr1[2])

	// Slices
	fmt.Println(arr1[2:])
	fmt.Println(arr1[2:3])
	fmt.Println(arr1[:2])
	fmt.Println(arr1[:])

	// Length
	fmt.Println(len(arr1))

	// Assign
	arr1[2] = 10
	fmt.Println(arr1)

	// Capacity
	fmt.Println(cap(arr1)) // capacity is the maximum number of elements that can be stored in the array

	// Append
	arr3 := append(arr1, 10, 11, 12) // append will create a new array with a larger capacity if needed, but does not change the original array
	// arr2 = append(arr2, 6) // This will not work as arr2 is a fixed size array
	fmt.Println(arr3)

	// Desestructuring or unpacking
	fmt.Println(append(arr1, arr2[:]...)) // This will append all elements of arr2 to arr1 as arr2 is a fixed size array
	fmt.Println(append(arr1, arr3...))    // This will append all elements of arr3 to arr1 as arr3 is a slice

	// Special function make to create a slice -> when we know roughly how many elements we want to store
	// This is more efficient than creating a slice with no size and then adding elements to it
	arr4 := make([]int, 5)     // This will create a slice of size 5 with all elements initialized to 0
	arr5 := make([]int, 5, 10) // This will create a slice of size 5 with all elements initialized to 0 and a capacity of 10
	fmt.Println(arr4)
	fmt.Println(arr5)

	// Iterate over a slice
	for index, value := range arr1 {
		fmt.Println(concatenate("Index: ", index), "-", concatenate("Value: ", value))
	}
}

func concatenate(values ...interface{}) string {
	var result string
	for _, value := range values {
		if str, ok := value.(string); ok {
			result += str
		} else if num, ok := value.(int); ok {
			result += fmt.Sprintf("%d", num)
		} else if flt, ok := value.(float64); ok {
			result += fmt.Sprintf("%f", flt)
		} else {
			result += fmt.Sprintf("%v", value)
		}
	}
	return result
}

// go run app.go

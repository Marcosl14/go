package main

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Using the transformNumbers function with different operations
	squared := transformNumbers(&numbers, square)
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	println("Squared:", squared)
	println("Doubled:", doubled)
	println("Tripled:", tripled)
}

// This function takes a slice of integers and a function as parameters.
func transformNumbers(numbers *[]int, operation func(int) int) []int {
	var result []int
	for _, number := range *numbers {
		result = append(result, operation(number))
	}
	return result
}

func square(n int) int {
	return n * n
}
func double(n int) int {
	return n * 2
}
func triple(n int) int {
	return n * 3
}

// go run app.go

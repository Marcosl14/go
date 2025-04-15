package main

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Closure functions for different operations
	square := createMultiplierFunction(2)
	double := createMultiplierFunction(2)
	triple := createMultiplierFunction(3)

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

// This function creates a closure that multiplies a number by a given factor.
// Closures are functions that capture the environment in which they were created.
// In this case, the closure captures the 'factor' variable.
// The returned function can be called with a number to multiply it by the factor.
// This is a common pattern in functional programming.
func createMultiplierFunction(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// go run app.go

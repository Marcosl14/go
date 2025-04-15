package main

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Anonymous functions for different operations
	square := func(n int) int {
		return n * n
	}
	double := func(n int) int {
		return n * 2
	}

	squared := transformNumbers(&numbers, square)
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, func(n int) int {
		return n * 3
	})

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

// go run app.go

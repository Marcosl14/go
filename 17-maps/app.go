package main

import "fmt"

// Type aliases
type floatMap map[string]float64

func main() {
	// Create a map -> key-value pairs
	// The key is a string and the value is an int
	ages := map[string]int{
		"John": 30,
		"Jane": 25,
	}

	// Add a new key-value pair
	ages["Bob"] = 40

	// Update an existing key-value pair
	ages["Jane"] = 26

	// Delete a key-value pair
	delete(ages, "John")

	// Print the map
	for name, age := range ages {
		fmt.Println(name, age)
	}

	// Check if a key exists
	if age, exists := ages["Bob"]; exists {
		fmt.Println("Bob's age is", age)
	} else {
		fmt.Println("Bob not found")
	}

	// Get the length of the map
	fmt.Println("Number of entries in the map:", len(ages))

	// Iterate over the map
	for name, age := range ages {
		fmt.Println(name, age)
	}

	// Special function called make, to create a map -> when we know roughly how many elements we want to store
	// This is more efficient than creating a map with no size and then adding elements to it
	var ages2 = make(map[string]int, 3)
	ages2["Alice"] = 35
	ages2["Charlie"] = 28
	ages2["Dave"] = 22
	ages2["Eve"] = 29 // This will work but it will not be efficient
	fmt.Println("Number of entries in the ages map:", len(ages2))

	var lengths = make(floatMap, 3)
	lengths["circle"] = 3.14
	lengths["square"] = 4.0
	lengths["triangle"] = 5.0
	fmt.Println("Number of entries in the length map:", len(lengths))

	// Iterate over the map
	for key, value := range lengths {
		fmt.Println(concatenate("Key: ", key), "-", concatenate("Value: ", value))
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

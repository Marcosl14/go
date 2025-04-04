package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	writeFile("Hello, World!")

	data, err := readFile()

	if err != nil {
		fmt.Println("Error:", err)
		panic("Can't continue") // This will stop the program
	}

	fmt.Println("Data from file:", data)
}

func writeFile(file string) {
	os.WriteFile("../00-others/output.txt", []byte(file), 0644)
	fmt.Println("File written successfully")
}

func readFile() (string, error) {
	data, err := os.ReadFile("../00-others/output.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", errors.New("file not found")
	}

	fmt.Println("File read successfully")
	return string(data), nil
}

// go run app.go

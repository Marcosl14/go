package main

import (
	"fmt"
	"time"
)

func main() {
	greet()

	// Asynchronous programming in Go
	// 1. Create a channel
	// 2. Create a goroutine that sends data to the channel
	// 3. Read the data from the channel in the main function
	// 4. Close the channel

	ch := make(chan bool)
	go slowGreet(ch)
	// msg := <-ch
	// fmt.Println(msg)
	<-ch
	close(ch)

	// Slice example
	var numbers []int = []int{1, 2, 3, 4, 5}
	var result []int

	doneChannels := make([]chan int, len(numbers))

	for index, number := range numbers {
		doneChannels[index] = make(chan int)
		go square(number, doneChannels[index])
		result = append(result, <-doneChannels[index])
	}

	fmt.Println(result)

	// Error handling
	errorChannel := make(chan error)
	doneChannel := make(chan bool)
	go someFunctionWithErrorHandling(doneChannel, errorChannel)
	select {
	case err := <-errorChannel:
		fmt.Println("Error:", err)
	case <-doneChannel:
		fmt.Println("Done")
	}

	// Defer
	defer func() {
		fmt.Println("Defer will be executed at the end of the higher scope function")
	}()

	greet()
	// Defer will execute after greet
}

func greet() {
	fmt.Println("Hello, World!")
}

func slowGreet(channel chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello, Slow World!")
	channel <- true
}

func square(n int, channel chan int) {
	channel <- n * n
}

func someFunctionWithErrorHandling(channel chan bool, errorChannel chan error) {
	// Simulate some work
	time.Sleep(2 * time.Second)

	// Simulate an error
	if true {
		errorChannel <- fmt.Errorf("an error occurred")
		return
	}

	channel <- true
}

// go run app.go

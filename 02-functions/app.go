package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(1, 2))

	something1, something2 := smth1()
	fmt.Println(something1, something2)

	fr1, fr2 := smth2()
	fmt.Println(fr1, fr2)

	val1, val2 := smth2()
	fmt.Println(val1, val2)
}

func add(x int, y int) int {
	return x + y
}

func smth1() (int, int) {
	return 1, 2
}

func smth2() (fr1 int, fr2 int) {
	fr1 = 1
	fr2 = 2
	// return fr1, fr2
	return
}

// go run app.go

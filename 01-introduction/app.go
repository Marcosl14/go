package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello Wolrd")

	var numVal1 = 1000
	var numVal2 = 2000

	fmt.Println("numVal1:", numVal1)
	fmt.Println("numVal2:", numVal2)
	fmt.Println("numVal1 + numVal2:", numVal1+numVal2)
	fmt.Println("numVal1 - numVal2:", numVal1-numVal2)
	fmt.Println("numVal1 * numVal2:", numVal1*numVal2)
	fmt.Println("numVal1 / numVal2:", numVal1/numVal2)
	fmt.Println("numVal1 % numVal2:", numVal1%numVal2)

	fmt.Println("numVal1 == numVal2:", numVal1==numVal2)
	fmt.Println("numVal1 != numVal2:", numVal1!=numVal2)
	fmt.Println("numVal1 > numVal2:", numVal1>numVal2)
	fmt.Println("numVal1 < numVal2:", numVal1<numVal2)
	fmt.Println("numVal1 >= numVal2:", numVal1>=numVal2)
	fmt.Println("numVal1 <= numVal2:", numVal1<=numVal2)

	fmt.Println("numVal1 & numVal2:", numVal1&numVal2) // AND
	fmt.Println("numVal1 | numVal2:", numVal1|numVal2) // OR
	fmt.Println("numVal1 ^ numVal2:", numVal1^numVal2) // XOR
	fmt.Println("numVal1 &^ numVal2:", numVal1&^numVal2) // AND NOT
	fmt.Println("numVal1 << numVal2:", numVal1<<numVal2) // left shift
	fmt.Println("numVal1 >> numVal2:", numVal1>>numVal2) // right shift

	fmt.Println("------------------------")
	fmt.Println("------------------------")

	var numVal3 = 100
	var numVal4 = 0.5

	fmt.Println("numVal3:", numVal3)
	fmt.Println("numVal4:", numVal4)

	var multiplication = float64(numVal3) * numVal4
	fmt.Println("numVal3 * numVal4:", multiplication)

	var power = math.Pow(float64(numVal3), numVal4)
	fmt.Println("numVal3 ^ numVal4:", power)

	fmt.Println("------------------------")
	fmt.Println("------------------------")

	var numVal5 float64 = 100
	var numVal6 float64 = 0.2
	var multiplication2 = numVal5 * numVal6 // no need to convert
	fmt.Println("numVal5 * numVal6:", multiplication2)
	var numVal7 int = int(numVal6)
	fmt.Println("numVal7:", numVal7)

	fmt.Println("------------------------")
	fmt.Println("------------------------")

	numVal8 := 100 // short declaration (can't type)
	var numVal9, numVal10 int = 200, 300 // multiple declaration same type (forced because of int)
	var numVal11, numVal12 = 200, "300" // multiple declaration different type
	numVal13, numVal14 := 200, "300" // multiple declaration different type
	fmt.Println("numVal8:", numVal8)
	fmt.Println("numVal9:", numVal9)
	fmt.Println("numVal10:", numVal10)
	fmt.Println("numVal11:", numVal11)
	fmt.Println("numVal12:", numVal12)
	fmt.Println("numVal13:", numVal13)
	fmt.Println("numVal14:", numVal14)

	fmt.Println("------------------------")
	fmt.Println("------------------------")

	const numVal15 = 100
	fmt.Println("const numVal15:", numVal15)

	fmt.Println("------------------------")
	fmt.Println("------------------------")

	var userInput string
	fmt.Print("Enter your name: ")
	fmt.Scan(&userInput) // read input from user -> & allows to change the value of userInput
	fmt.Println("Your name is:", userInput)
}

// go run app.go

package main

import (
	"fmt"

	"example.com/module/otherPackages"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	externalFunction()                                                                 // same package functions are not neccesarily capitalized
	otherPackages.ExternalFunction()                                                   // other packages function names must be capitalized
	fmt.Println("Third Party Library:", randomdata.FirstName(randomdata.RandomGender)) // third party library
}

// external functions only works in modules!!!
// go mod init example.com/module
// go run .

// installing third party libraries
// go get github.com/Pallinder/go-randomdata

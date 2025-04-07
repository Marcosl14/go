package main

import (
	"fmt"
	"time"
)

func main() {
	// Structs are used to group related data together
	// Struct Literal or Collection Literal

	var person1 Person = Person{
		Name:      "John Doe",
		Age:       30,
		CreatedAt: time.Now(),
	}

	fmt.Println("Person 1:", person1)

	var person2 Person = Person{
		"Jane Doe",
		25,
		time.Now(),
	}

	fmt.Println("Person 2:", person2)

	fmt.Println("------------------------")
	fmt.Println("------------------------")

	fmt.Println("Null Value")
	var person3 Person
	fmt.Println("Person 3:", person3)

	fmt.Println("------------------------")
	fmt.Println("------------------------")

	fmt.Println("Structs in Regular Functions")
	fmt.Println("Person 1 Pointer:", &person1)
	fmt.Println("Person 1 Name:", person1.Name)
	fmt.Println("Person 1 Name Pointer:", &person1.Name)
	printPersonCopy(person1)

	fmt.Println("------------------------")
	fmt.Println("------------------------")

	fmt.Println("Structs in Pointer Functions")
	fmt.Println("Person 1 Pointer:", &person1)
	fmt.Println("Person 1 Name:", person1.Name)
	fmt.Println("Person 1 Name Pointer:", &person1.Name)
	printSamePerson(&person1)

	fmt.Println("------------------------")
	fmt.Println("------------------------")

	fmt.Println("Struct Methods")
	person1.Print()               // call the method on the struct
	person1.SetName("Luis Smith") // call the method on the struct
	person1.Print()               // call the method on the struct

	fmt.Println("------------------------")
	fmt.Println("------------------------")

	fmt.Println("Constructor Function")
	newPerson("John Smith", 40).Print() // call the function that returns a struct and call the method on it
}

func printPersonCopy(person Person) {
	fmt.Println("Function Person Pointer:", &person)
	fmt.Println("Function Person Name:", person.Name)
	fmt.Println("Function Person Name Pointer:", &person.Name) // its a copy of the original value
}

func printSamePerson(person *Person) {
	fmt.Println("Function Person Pointer:", &person)
	fmt.Println("Function Person Name:", person.Name)
	fmt.Println("Function Person Name Pointer:", &person.Name) // its the original value (same memory address)
}

// --------------------------------------------------------
// Struct Person

type Person struct {
	Name      string
	Age       int
	CreatedAt time.Time
}

// Struct methods are functions that are associated with a struct by calling the struct before the function name ( (p Person) )
func (p Person) Print() { // p is a copy of the struct
	fmt.Println("Person Name:", p.Name)
	fmt.Println("Person Age:", p.Age)
	fmt.Println("Person CreatedAt:", p.CreatedAt)
}

func (p *Person) SetName(name string) { // p is a pointer to the struct
	p.Name = name
}

func newPerson(name string, age int) Person {
	return Person{
		Name:      name,
		Age:       age,
		CreatedAt: time.Now(),
	}
}

// Struct Person
// --------------------------------------------------------

// go run app.go

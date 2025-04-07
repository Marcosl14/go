package main

import (
	"fmt"

	"example.com/structs/person"
)

func main() {
	var person1 person.Person = person.Person{
		// name: "John Doe", // don't have access
		// age:  30, // don't have access
	}
	fmt.Println("Person 1:", person1)

	// instead use:
	var person2 *person.Person = person.NewPerson("John Doe", 30)
	fmt.Println("Person 2:", person2)

	fmt.Println("------------------------")
	fmt.Println("------------------------")

	var employee1 person.Employee = *person.NewEmployee(person2, 1000)
	fmt.Println("Employee 1:", employee1)
	// employee1.Person.SetName("Luis Smith") // I can't access because person is not exported

	fmt.Println("------------------------")
	fmt.Println("------------------------")

	var student1 person.Student = *person.NewStudent("John Doe", 20, 1000)
	fmt.Println("Student 1:", student1)
	student1.SetName("Smith Jackson") // I can access because person is exported
	student1.Person.Print()           // I can access this way too
}

// go run app.go

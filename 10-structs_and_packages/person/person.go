package person

import (
	"fmt"
	"time"
)

// All that has to be exported must be capitalized even struct name, fields and methods
// What is not capitalized is not exported
type Person struct {
	name      string    `json:"name"`       // json tag (metadata)
	age       int       `json:"age"`        // json tag (metadata)
	createdAt time.Time `json:"created_at"` // json tag (metadata)
}

// Struct methods are functions that are associated with a struct by calling the struct before the function name ( (p Person) )
func (p Person) Print() { // p is a copy of the struct
	fmt.Println("Person Name:", p.name)
	fmt.Println("Person Age:", p.age)
	fmt.Println("Person CreatedAt:", p.createdAt)
}

func (p *Person) SetName(name string) { // p is a pointer to the struct
	p.name = name
}

func NewPerson(name string, age int) *Person {
	return &Person{
		name:      name,
		age:       age,
		createdAt: time.Now(),
	}
}

// inheritance

type Employee struct {
	person Person
	salary int
}

type Student struct {
	Person
	average int
}

func NewEmployee(person *Person, salary int) *Employee {
	return &Employee{
		person: *person,
		salary: salary,
	}
}

func NewEmployee2(name string, age int, salary int) *Employee {
	return &Employee{
		person: Person{
			name: name,
			age:  age,
		},
		salary: salary,
	}
}

func NewEmployee3(name string, age int, salary int) *Employee {
	return &Employee{
		person: *NewPerson(name, age),
		salary: salary,
	}
}

func NewStudent(name string, age int, average int) *Student {
	return &Student{
		Person:  *NewPerson(name, age),
		average: average,
	}
}

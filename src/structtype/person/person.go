package person

import (
	"fmt"
)

type Person struct {
	Name    string
	Address string
	Phone   int64
}

// returns a pointer to avoid copying the structure
func CreatePerson() *Person {
	var newPerson Person
	newPerson.Name = "mario"
	newPerson.Address = "rome"
	newPerson.Phone = 349121454
	return &newPerson
}

// receives a pointer to avoid copying the structure
func ShowPerson(aPerson *Person) {
	fmt.Printf("Name: %s\n", aPerson.Name)
	fmt.Printf("Address: %s\n", aPerson.Address)
	fmt.Printf("Phone: %d\n", aPerson.Phone)
}

// function receives a pointer, because structures are copied by value
func ChangePersonPhone(aPerson *Person, newPhone int64) {
	aPerson.Phone = newPhone
}

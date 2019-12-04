package main

import (
	"fmt"
	"structtype/person"
)

func main() {

	// use a type with struct
	var person1 person.Person
	person1.Name = "john"
	person1.Address = "lisbon"
	person1.Phone = 959933303
	fmt.Println(person1)

	// type struct literal
	person2 := person.Person{Name: "mary", Address: "madrid", Phone: 878434421}
	fmt.Println(person2)

	// create a person and return the data as a type
	person3 := person.CreatePerson()
	fmt.Println(person3)

	// send a person struct type to a function
	person.ShowPerson(person3)

	// send a person struct type to a function
	// as person1 is a variable, a reference is necessary to have
	// a pointer in the function
	person.ChangePersonPhone(&person1, 667243434)
	person.ShowPerson(&person1)

	// as person3 is a pointer, the pointer value is used directly
	person.ChangePersonPhone(person3, 234323562)
	person.ShowPerson(person3)
}

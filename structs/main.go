package main

import "fmt"

type person struct {
	firstName string
	lastName string
	contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Boe",
		contactInfo: contactInfo{
			email: "jim@hotmail.com",
			zipCode: 90210,
		},
	}

	jim.updateName("jimmy")
	jim.printPerson()
}

func (p person) printPerson() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
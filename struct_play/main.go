package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// var alex person

	// alex.firstName = "test"
	// alex.lastName = "no test"

	// fmt.Printf("%+v", alex)
	jim := person{
		firstName: "Jim",
		lastName:  "pretty",
		contactInfo: contactInfo{
			email:   "test",
			zipCode: 123,
		},
	}

	fmt.Println(jim.contactInfo.email)

	jim.updateName("Test")
	jim.print()

}

func (pointerToPerson *person) updateName(newFieldName string) {
	(*pointerToPerson).firstName = newFieldName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

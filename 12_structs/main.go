package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Value receiver method: Greeting
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " +
		person.lastName + ". I am " + strconv.Itoa(person.age) +
		" and I come from " + person.city
}

// Pointer receiver method: hasBirthday
func (person *Person) hasBirthday() {
	person.age++
}

// Pointer receiver method: marriedTo
func (person *Person) marriedTo(anotherPerson Person) {
	if person.gender == "M" {
		return
	}

	person.lastName = anotherPerson.lastName
}

func main() {
	// Init person
	alex := Person{
		firstName: "Alex",
		lastName:  "Such",
	}

	alex.age = 37
	alex.city = "Alicante"
	alex.gender = "M"

	sandra := Person{"Sandra", "Lorente", "Alicante", "F", 35}

	fmt.Println(alex)
	fmt.Println(sandra)

	fmt.Println(sandra.greet())
	fmt.Println(alex.greet())

	alex.hasBirthday()
	fmt.Println(alex.greet())

	alex.marriedTo(sandra)
	sandra.marriedTo(alex)
	fmt.Println(alex.greet())
	fmt.Println(sandra.greet())

}

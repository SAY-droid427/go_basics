package main

import (
	"fmt"
	"strconv"
)

// Person struct...
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	// Shorter way
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method(value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + "and I am " + strconv.Itoa(p.age)
}

// hasBirthday method(pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried(pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {

	} else {
		p.lastName = spouseLastName
	}
}
func main() {
	//Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "London", gender: "f", age: 25}
	fmt.Println(person1)
	// Alternative
	person2 := Person{"Samantha", "Smith", "London", "f", 25}
	fmt.Println(person2)

	person1.hasBirthday()
	fmt.Println(person1.greet())
	person1.getMarried("Johnson")
}

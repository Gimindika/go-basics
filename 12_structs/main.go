package main

import (
	"fmt"
	"strconv"
)

//defining struct
type Person struct {
	//declare struct properties
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	//cleaner way to declare struct properties
	firstName, lastName, city, gender string
	age                               int
}

// struct method (value reciever) - doesnt change anything
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and i am " + strconv.Itoa(p.age)
}

// struct method (pointer reciever) - change something
func (p *Person) hasBirthday() {
	p.age++
}

//another example - getMarried(pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	//initialize struct
	person1 := Person{firstName: "Gerrit", lastName: "Indika", city: "Plb", gender: "M", age: 25}
	fmt.Println(person1)

	person2 := Person{"Hana", "Sharon", "Plb", "F", 18}
	fmt.Println(person2)

	// fmt.Println(person1.firstName)

	// person1.hasBirthday()

	// fmt.Println(person1.greet())

	person3 := Person{"Yosua", "Wewengkang", "Plb", "M", 22}
	person2.getMarried(person3.lastName)
	fmt.Println(person2.greet())

}

package main

import "fmt"

func main() {
	var isYes bool
	isYes = true

	//constant
	const name string = "Gerrit"

	//omit type
	var car = "Hondi"

	//shorthand - inside function only
	age := 31
	pet, color := "dog", "blue"

	//declaring variable without using it will resulting in error
	fmt.Println(name, car, age, pet, color)

	//type check
	fmt.Printf("%T", isYes)
}

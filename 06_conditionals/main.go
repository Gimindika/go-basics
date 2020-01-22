package main

import "fmt"

func main() {
	x := 15
	y := 10

	if x <= y {
		fmt.Printf("%d is less than or equal to %d", x, y)
	} else {
		fmt.Printf("%d is less than %d", y, x)
	}

	fmt.Println("")

	color := "red"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is not blue or red")

	}

	pet := "Dog"
	switch pet {
	case "Dog":
		fmt.Println("Pet is Dog")

	case "Cat":
		fmt.Println("Pet is Cat")

	default:
		fmt.Println("Pet is not Dog or Cat")
	}
}

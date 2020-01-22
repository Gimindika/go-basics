package main

import "fmt"

func main() {
	a := 5
	b := &a //pointer to memmory address/locaction of variable a

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	fmt.Println("")
	//reading value from memmory address
	fmt.Println(b)
	fmt.Println(*b)

	fmt.Println("")
	//change value with pointer
	fmt.Println(a)
	*b = 10
	fmt.Println(a)
}

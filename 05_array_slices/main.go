package main

import "fmt"

func main() {
	var fruits [2]string
	fruits[0] = "Melon"
	fruits[1] = "Grape"
	fmt.Println(fruits)

	pets := [3]string{"Dog", "Cat", "Bunny"}
	fmt.Println(pets)

	petSlice := []string{"Dog", "Pig", "Bunny", "Bear"}
	fmt.Println(petSlice)
	fmt.Println(len(petSlice))
	fmt.Println(petSlice[1:3])
}

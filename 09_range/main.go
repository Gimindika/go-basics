package main

import "fmt"

func main() {
	// range with array
	ids := []int{12, 32, 43, 33, 54}

	//with index
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//without index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//sum
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum ", sum)

	//range with map -> k = key, v = value
	emails := map[string]string{"Gerrit": "gerrit@gmail.com", "Bob": "bob@gmail.com", "Paulina": "paulina@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	//key only
	for k := range emails {
		fmt.Println("Name : " + k)
	}

	//value only
	for _, v := range emails {
		fmt.Println("Email : ", v)
	}
}

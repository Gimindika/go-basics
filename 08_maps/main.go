package main

import "fmt"

func main() {
	// map = key - value pair
	emails := make(map[string]string)
	emails["Gerrit"] = "gimindika@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Sharon"])

	delete(emails, "Mike")
	fmt.Println(emails)

	emailsAgain := map[string]string{"Bob": "bob@gmail.com", "Paulina": "paulina@gmail.com"}
	fmt.Println(emailsAgain)
}

package main

import (
	"fmt"
	"math"

	"github.com/Gimindika/basics/03_packages/strutil" //local custom package
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("hello"))
}

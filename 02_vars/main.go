package main

import (
	"fmt"
)

func main() {
	// Declaring var
	var name = "Alejandro"
	var age uint8 = 37
	isCool := true

	fmt.Printf("Hello %s. I heard you are %d, %T\n", name, age, isCool)

	fmt.Printf("%s is a %T\n", name, name)
	fmt.Printf("%d is a %T\n", age, age)
}

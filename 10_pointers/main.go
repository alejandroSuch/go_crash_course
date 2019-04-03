package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Printf("%d - %d (%d)\n", a, *b, b)

	a = 4
	fmt.Printf("%d - %d (%d)\n", a, *b, b)

	*b = 14
	fmt.Printf("%d - %d (%d)\n", a, *b, b)

	*&a = 15
	fmt.Printf("%d - %d (%d)\n", a, *b, b)
}

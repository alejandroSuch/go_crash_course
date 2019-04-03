package main

import "fmt"

func main() {
	x := 10
	y := 10

	if x < y {
		fmt.Printf("%d is lower than %d\n\n", x, y)
	} else {
		fmt.Printf("%d is greater than or equals %d\n\n", x, y)
	}

	color := "orange"

	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not red nor blue")

	}
}

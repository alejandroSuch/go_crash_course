package main

import "fmt"

func main() {
	// Arrays --> fixed length
	var fruits [2]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"

	fmt.Println(fruits)
	fmt.Println(fruits[0])
	fmt.Println(fruits[1])

	moreFruits := [2]string{"Apple2", "Banana2"}
	fmt.Println(moreFruits)

	// Slices --> uncertain length
	fruitsSlice := []string{"Pear", "Strawberry", "Grape"}
	fruitsSlice[1] = "Cherry"

	fmt.Println(fruitsSlice)
	fmt.Println(fruitsSlice[1:3])
	fmt.Println(len(fruitsSlice))
}

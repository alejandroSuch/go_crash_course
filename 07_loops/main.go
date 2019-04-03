package main

import (
	"fmt"
)

func main() {
	// Long method
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	// Short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Counter is %d\n`", i+1)
	}

	// FizzBuzz
	for i := 1; i < 100; i++ {
		fmt.Printf("%d ", i)

		if i%3 == 0 {
			fmt.Print("Fizz")
		}

		if i%5 == 0 {
			fmt.Print("Buzz")
		}

		fmt.Println()
	}
}

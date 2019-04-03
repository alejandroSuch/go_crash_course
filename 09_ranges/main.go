package main

import "fmt"

func main() {
	ids := []int{2, 34, 67, 75, 12}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID %d\n", i, id)
	}

	// Loop through ids not using index
	for _, id := range ids {
		fmt.Printf("ID %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Printf("Sum is %d\n", sum)

	// Range in maps
	emails := map[string]string{
		"alex":   "alex@gmail.com",
		"sandra": "sandra@gmail.com",
	}

	for name, email := range emails {
		fmt.Printf("%s - %s\n", name, email)
	}

	for name := range emails {
		fmt.Printf("%s: %s\n", name, emails[name])
	}
}

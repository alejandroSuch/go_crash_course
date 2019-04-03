package main

import (
	"fmt"
)

func main() {
	// Define maps
	emails := make(map[string]string)

	// Assign keys and values
	emails["alex"] = "alex@gmail.com"
	emails["sandra"] = "sandra@gmail.com"
	emails["martina"] = "martina@gmail.com"
	emails["mario"] = "mario@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["alex"])

	// Delete one
	delete(emails, "alex")
	emails["yayo"] = "yayo@gmail.com"
	fmt.Println(emails)

	// Do not make
	moreEmails := map[string]string{
		"alex":   "alex@gmail.com",
		"sandra": "sandra@gmail.com",
	}

	moreEmails["mario"] = "mario@gmail.com"
	delete(moreEmails, "alex")

	fmt.Println(moreEmails)

	anotherEmailList := map[string]string{}
	fmt.Println(anotherEmailList)
	anotherEmailList["martina"] = "martina@gmail.com"
	fmt.Println(anotherEmailList)
}

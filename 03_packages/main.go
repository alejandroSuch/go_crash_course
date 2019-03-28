package main

import (
	"fmt"
	"math"
	"tutorial/03_packages/strutil"
)

func main() {
	fmt.Printf("Hello, %f\n", math.Floor(1.5))

	fmt.Printf("Hola al revés es %s\n", strutil.Reverse("hola"))
}

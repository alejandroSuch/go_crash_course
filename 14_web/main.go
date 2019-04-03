package main

import (
	"fmt"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>Hello, World</h1>")
}

func about(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>About</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("Server running at port 3000")
	http.ListenAndServe(":3000", nil)
}

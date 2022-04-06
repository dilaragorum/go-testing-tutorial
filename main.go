package main

import (
	"fmt"
	"net/http"
)

func Greeting() string {
	return "Hello"
}

func GoodBye(name string) string {
	text := "Goodbye, " + name + "..."
	return text
}

func Add(a int, b int) int {
	return a + b
}

func main() {
	http.HandleFunc("/", GreetingAPI)
	http.ListenAndServe(":3000", nil)
}

func GreetingAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gopher")
}

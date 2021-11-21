package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, input string) {
	fmt.Fprintf(writer, "Hello, %s", input)
}

// http.ResponseWriter implements io.Writer
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Elodie")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}

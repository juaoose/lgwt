package main

import "fmt"

func Hello(recipient string) string {
	return "Hello, " + recipient
}

func main() {
	fmt.Println(Hello("Juan"))
}

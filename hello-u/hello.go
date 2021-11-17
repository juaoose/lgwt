package main

import "fmt"

const helloPrefix = "Hello, "

func Hello(recipient string) string {
	return helloPrefix + recipient
}

func main() {
	fmt.Println(Hello("Juan"))
}

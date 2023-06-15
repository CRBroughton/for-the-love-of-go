package main

import "fmt"

func mySwitch(name string) string {
	switch name {
	case "Alice":
		return "Hey, Alice."
	case "Bob":

		return "What's up, Bob."
	default:
		return "Hello, stranger."
	}
}

func main() {
	value := mySwitch("Alice")

	fmt.Println(value)
}

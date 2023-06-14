package main

import "fmt"

func bigger(first, second int) int {
	if first > second {
		return first
	}
	return second
}

func main() {
	value := bigger(4, 5)

	fmt.Println(value)
}

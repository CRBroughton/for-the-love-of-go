package main

import "fmt"

func apply(value int, f func(x int) int) int {
	return f(value)
}

func main() {
	value := apply(1, func(x int) int {
		return x * 2
	})

	fmt.Println(value)
}

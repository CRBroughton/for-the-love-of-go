package main

import "fmt"

func total(values []int) int {
	var value int

	for _, e := range values {
		value = value + e
	}
	return value
}

func main() {
	value := total([]int{1, 2, 3})

	fmt.Println(value)
}

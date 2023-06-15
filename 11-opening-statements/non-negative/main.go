package main

import "fmt"

func nonNegative(values []int) []int {
	var value []int

	for _, e := range values {
		if e >= 0 {
			value = append(value, e)
		}
	}
	return value
}

func main() {
	value := nonNegative([]int{2, -1, 0, -4})

	fmt.Println(value)
}

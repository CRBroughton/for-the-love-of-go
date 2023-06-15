package main

import "fmt"

func evens() []int {
	var values []int

	for x := 0; x < 100; x++ {
		if x%2 == 0 {
			values = append(values, x)
		}
	}
	return values
}

func main() {
	value := evens()

	fmt.Println(value)
}

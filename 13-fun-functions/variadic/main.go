package main

import "fmt"

func AddMany(inputs ...float64) float64 {
	var value float64
	for _, input := range inputs {
		value = value + input
	}
	return value
}

func main() {
	value := AddMany(1, 2, 3, 4, 5)

	fmt.Println(value)
}

package main

import "fmt"

func xor(first, second bool) bool {
	if first == true || second == true {
		return true
	}
	return false
}

func main() {
	result := xor(true, false)

	fmt.Println(result)
}

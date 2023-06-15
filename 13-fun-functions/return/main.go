package main

import "fmt"

func withdraw(balance, amount int) (int, error) {
	var newBalance int
	if amount > balance {
		return 0, fmt.Errorf("nope")
	}
	newBalance = balance - amount
	return newBalance, nil
}

func main() {
	value, err := withdraw(2, 4)

	if err != nil {
		fmt.Print("You cannot draw this amount")
	}

	fmt.Println(value)
}

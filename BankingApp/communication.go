package main

import "fmt"

// Showing example of splitting code into multiple files across the same pacakage main
func presentOptions() {
	fmt.Println("What do you want to do today?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName string = "dotmanBalance.txt"

func getBalanceFromFile() (float64, error) {
	// Typically errors doesn't crash proejct in Go, rather returns a default value
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse srtored balance value")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	// 0644 is a file permission notation for operating systems(especially) that allows owner to read and write but others can only read it
	// https://www.redhat.com/sysadmin/linux-file-permissions-explained
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		return
		//Panic  will end the application and also return the error message
		// panic("Can't continue sorry")
	}

	fmt.Println("Account Balance :", accountBalance)
	fmt.Println("Welcome to Dotman Go Bank")

	for {
		fmt.Println("What do you want to do today?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("How much do you wanna deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0")
				continue //This will restart the code from the start again
			}
			// accountBalance = accountBalance + depositAmount
			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance Updated, new amount: ", accountBalance)
		case 3:
			var withdrawalAmount float64
			fmt.Print("How much do you wanna withdraw: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid Amount. You can't withdraw more than you have")
				continue
			}

			accountBalance -= withdrawalAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance Updated, new amount: ", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for banking with Dotman")
			return
		}
	}
}

// func exampleWIthIfElse() {

// 	var accountBalance float64 = 1000.0

// 	fmt.Println("Welcome to Dotman Go Bank")
// for i := 0; i < 2; i++ {  A normal for loop
// This will run an infinite loop
// 	for {
// 		fmt.Println("What do you want to do today?")
// 		fmt.Println("1. Check balance")
// 		fmt.Println("2. Deposit money")
// 		fmt.Println("3. Withdraw money")
// 		fmt.Println("4. Exit")

// 		var choice int
// 		fmt.Print("Your choice: ")
// 		fmt.Scan(&choice)

// 		if choice == 1 {
// 			fmt.Println("Your balance is:", accountBalance)
// 		} else if choice == 2 {
// 			var depositAmount float64
// 			fmt.Print("How much do you wanna deposit: ")
// 			fmt.Scan(&depositAmount)
// 			if depositAmount <= 0 {
// 				fmt.Println("Invalid Amount. Must be greater than 0")
// 				continue //This will restart the code from the start again
// 			}
// 			// accountBalance = accountBalance + depositAmount
// 			accountBalance += depositAmount
// 			fmt.Println("Balance Updated, new amount: ", accountBalance)
// 		} else if choice == 3 {
// 			var withdrawalAmount float64
// 			fmt.Print("How much do you wanna withdraw: ")
// 			fmt.Scan(&withdrawalAmount)

// 			if withdrawalAmount <= 0 {
// 				fmt.Println("Invalid Amount. Must be greater than 0")
// 				continue
// 			}

// 			if withdrawalAmount > accountBalance {
// 				fmt.Println("Invalid Amount. You can't withdraw more than you have")
// 				continue
// 			}

// 			accountBalance -= withdrawalAmount
// 			fmt.Println("Balance Updated, new amount: ", accountBalance)
// 		} else {
// 			fmt.Println("Goodbye!")
// 			// return
// 			break // This breaks out of the loop
// 		}
// 	}
// 	fmt.Println("Thanks for banking with Dotman")
// }

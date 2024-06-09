package main

import (
	"fmt"

	"example.com/banking_app/fileOps"
	"github.com/Pallinder/go-randomdata"
)

const fileName string = "dotmanBalance.txt"

func main() {
	var accountBalance, err = fileOps.GetFloatFromFile(fileName)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		// return
		//Panic  will end the application and also return the error message
		// panic("Can't continue sorry")
	}

	fmt.Println("Account Balance :", accountBalance)
	fmt.Println("Welcome to Dotman Go Bank")

	// Showing the use of function from a package
	fmt.Println(randomdata.PhoneNumber())

	for {
		presentOptions()
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
			fileOps.WriteFloatToFile(accountBalance, fileName)
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
			fileOps.WriteFloatToFile(accountBalance, fileName)
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

package main

import "fmt"

func main() {

	revenue := getUserInput("Enter Revenue: ")
	expenses := getUserInput("Enter Expenses: ")
	taxRate := getUserInput("Enter Tax Rate: ")

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	ebt, profit, ratio := calculateFinacials(revenue, expenses, taxRate)

	fmt.Printf("Earning Before Tax: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.1f\n", ratio)
}

func getUserInput(inputText string) float64 {
	var userInput float64
	fmt.Print(inputText)
	fmt.Scan(&userInput)

	return userInput
}

func calculateFinacials(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

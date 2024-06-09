package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err1 := getUserInput("Enter Revenue: ")
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	expenses, err2 := getUserInput("Enter Expenses: ")
	if err2 != nil {
		fmt.Println(err1)
		return
	}
	taxRate, err3 := getUserInput("Enter Tax Rate: ")
	if err3 != nil {
		fmt.Println(err1)
		return
	}

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	// if err1 != nil || err2 != nil || err3 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	ebt, profit, ratio := calculateFinacials(revenue, expenses, taxRate)

	fmt.Printf("Earning Before Tax: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.1f\n", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\n Profit: %.1f\nRation: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func getUserInput(inputText string) (float64, error) {
	var userInput float64
	fmt.Print(inputText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("values must be a postive number")
	}

	return userInput, nil
}

func calculateFinacials(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

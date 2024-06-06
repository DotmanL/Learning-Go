// Must declare package main, each go project must have at least one package main.
// A project can have multiple packages and a single pacakge can be used in multiple files
package main

import (
	"fmt"
	"math"
)

// Basic Hello World Notes
// This initial function must be named main as it indicates the start of our application and needed by Go and must only be one even if we have
// multiple files with package main
// import "fmt"

// func main() {
// 	fmt.Print("Hello World")
// }

// To transform our project into a Go Module, run the command below in the terminal:
// - go mod init example.com/investment_calculator
// (ie: where investment_calculator is the app name and example.com is wherever with intend to host our module)
// This will generate a go.mod file.

// To build our project, just run the command below and this will generate a file with our app name(investment_calculator):
// - go build

// We can easily now run the module by just running below (ie: anyone can use our module even if they have no go installed):
// ./investment_calculator or by just running go .

const inflationRate = 6.5

func main() {
	var investmentAmount float64
	var years float64
	// The expectedReturnRate Variable declaration is the advised way of declaring variables with no explicit type but relying on the inferred type
	expectedReturnRate := 5.5

	// The & sign before a variable is used to declare a Pointer
	// fmt.Print("Enter Investment Amount: ")
	// Example using our own custom function
	outpuText("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter Expected  Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter Number of Years: ")
	fmt.Scan(&years)

	// This shows an example declaring multiple variables of the same type (Note: this can't be done with explicit different variable types)
	// var investMentAmount, years float64 = 1000, 10
	// Example showinng inferred types for multiple variables
	// investMentAmount, years := 1000, "10"

	// We can simplify our function to contain only one line if all varibles are decalred as float64
	// investmentAmount, years, expectedReturnRate := 1000.o, 10.0, 5.5

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, (years))
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	//Example using the function we created
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Real Value:", futureRealValue)

	// Printf use to format strings, eg: used to format above to two lines by just a one liner
	// %v  used to output a value in a default format
	// %f used to output a value in a default width, default precision
	// %.0f used to output a value with no decimal place as indicated by .0
	// %.1f used to output a value with one decimal place as indicated by .0
	// fmt.Printf("Future Value: %.1f\nFuture Real Value: %.1f", futureValue, futureRealValue)

	// Formatting using backticks can allow usn write a long line in multiple lines
	// fmt.Printf(`Future Value: %.1f

	// Future Real Value: %.1f`, futureValue, futureRealValue)

	// Sprintf formats according to a format specifier and returns the resulting string.
	fomattedFutureValue := fmt.Sprintf("Future Value %.1f\n", futureValue)
	fomattedRealFutureValue := fmt.Sprintf("Future Real Value (adjusted for inflation):  %.1f\n", futureRealValue)
	print(fomattedFutureValue, fomattedRealFutureValue)
}

// func outpuText(text string, text2 string) {
// Shorthand if parameters are of the same type
// func outpuText(text, text2 string) {
func outpuText(text string) {
	fmt.Print(text)

}

// You can return multiple values in a function in Go but must declare the return type
// func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
// 	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, (years))
// 	rfv := fv / math.Pow(1+inflationRate/100, years)
// 	return fv, rfv
// }

// Example showing using named returned types
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, (years))
	rfv = fv / math.Pow(1+inflationRate/100, years)
	// return fv, rfv
	return
}

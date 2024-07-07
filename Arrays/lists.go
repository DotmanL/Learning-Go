package main

import "fmt"

func main() {

	var productNames [4]string = [4]string{"A book"}
	//setting the third value
	productNames[2] = "A Carpet"
	// The type says we want to store 4 float64 values in this array
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
	fmt.Println(productNames)

	fmt.Println(prices[2])

	//This gives a slice of the array, a slice is just a tiny reference to an array and not a copy
	// featuredPrices := prices[1:3] // 9.99 45.99
	featuredPrices := prices[1:]
	featuredPrices[0] = 199.9
	highlightedPrices := featuredPrices[:1]
	// featuredPrices := prices[:3] starts from 0 index to the index 3 excluding
	// featuredPrices := prices[1:]  starts from 1 to the end
	fmt.Println(featuredPrices)
	fmt.Println(highlightedPrices)
	fmt.Println(len(featuredPrices), cap(featuredPrices)) //gives the length of the array an cap gives the capacity of the array

	highlightedPrices = highlightedPrices[:3] // this will pick more values to the right of the array as go doesn't memoize to the left of the initial highlightedPrices
	fmt.Println(featuredPrices)
	fmt.Println(highlightedPrices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))
}

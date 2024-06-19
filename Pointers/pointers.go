package main

import "fmt"

// Pointers are variables that store value addresses instead of values
// Why Pointers? 1. Avoid unnecessary value copies, 2. Directly Mutate values

// Its not worthy to use pointers with small data structures such as int, string etc.
// Not worth it in terms of conserving memory as they are small
func main() {

	age := 32 // regular variable

	// fmt.Println("Age:", age)
	// adultYears := getAdultYears(age)
	// fmt.Println(adultYears)

	// assuming we using a pointer
	// another way of declaring a pointer, the * before the type indicates its a pointer
	// var agePointer *int = &age

	// because of the & before the variable, we get the pointer address
	agePointer := &age

	// Avoid unnecessary value copies example
	// to get the value stored by the pointer(called derefrencing), use an * before the pointer variable like below
	// fmt.Println("Age Value:", *agePointer)
	// adultYears := getAdultYears(agePointer)
	// fmt.Println(adultYears)

	//just mutating the value example
	fmt.Println("Age Value:", *agePointer)
	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

// func getAdultYears(age int) int {
// 	return age - 18
// }

// Avoid unnecessary value copies example
// func getAdultYears(age *int) int {
// 	return *age - 18
// }

// Directly mutate the pointer value example
func editAgeToAdultYears(age *int) {
	*age = *age - 18
}

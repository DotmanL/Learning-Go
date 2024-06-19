package main

import (
	"fmt"

	"example.com/structs/user"
)

// Note: Check notes about Structs in the User Package, all code extracted there to demonstrate reusability

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// example by calling instance of the struct or a Struct Literal or Composite Literal
	//  var appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	// example using the normal creator function
	// var appUser = newUser(userFirstName, userLastName, userBirthdate)

	// example using the normal creator function with validation
	// var appUser *user
	// appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// appUser.outputUserDetails()
	// appUser.clearUserName()
	// appUser.outputUserDetails()

	//example using the user from the User pacakage

	var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("ola@gmail.com", "test123")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	usingCustomStypes()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

// Custom types
// the str we created is a custom type that still refers to the built in string type
type str string

func (text str) log() {
	fmt.Println(text)
}

func usingCustomStypes() {
	// Explicity define variable to use the custom type
	var name str = "Oladotun"

	name.log()
}

package main

import (
	"fmt"
	"time"
)

// A function attached to struct is called a method
type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// This function is now attached to the user struct by the prefix of (u user) -> Reciver argument before the function name
func (u user) outputUserDetails() {
	// No need to derefrence u ((*u).firstName) as with pointers and struct, the values can be accessed as below without the need of dereferencing
	fmt.Println(u.firstName, u.lastName, u.birthDate)

}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// instance of the struct or a Struct Literal or Composite Literal
	var appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

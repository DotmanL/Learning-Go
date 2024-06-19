package user

import (
	"errors"
	"fmt"
	"time"
)

// Non exported struct
// type user struct {
// Exported struct Only
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// Struct Embedding (equivalent of inheritance)
type Admin struct {
	email    string
	password string
	// User     User
	User // this is anonymous embedding, allows us to call methods on the embedded struct directly due to the anonymous embedding as opposed to above declaration
}

// We can export a struct by Capitalizing the first letter but not all the fields in the struct, to do that, need to
// capitalize all the fields
// type User struct {
// 	FirstName string
// 	lastName  string
// 	birthDate string
// 	createdAt time.Time
// }

// A function attached to struct is called a method
// This function is now attached to the user struct by the prefix of (u user) -> Reciver argument before the function name
func (u User) OutputUserDetails() {
	// No need to derefrence u ((*u).firstName) as with pointers and struct, the values can be accessed as below without the need of dereferencing
	fmt.Println(u.firstName, u.lastName, u.birthDate)

}

// When creating methods(ie: functios that have receiver arguments) to update/mutate a struct, pass the struct as the pointer value as below so to
// edit the original value stored in the address rather than copy. Same can be done for the outputUserDetails above but not mandatory since
// we are not updating the values.
func (u *User) ClearUserName() {
	u.firstName = " "
	u.lastName = " "
}

// using creator function for validations
func New(firstName, lastName, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

// Example showing embedded struct
func NewAdmin(email, password string) Admin {
	return Admin{email: email, password: password, User: User{
		firstName: "ADMIN",
		lastName:  "ADMIN",
		birthDate: "___",
		createdAt: time.Now(),
	}}
}

// Creator or Constructor function, not baked into go but mostly used as pattern, ie: a utility function that takes care of creating a struct
// start with new just a convention
// func newUser(firstName, lastName, birthDate string) user {
// 	return user{
// 		firstName: firstName,
// 		lastName:  lastName,
// 		birthDate: birthDate,
// 		createdAt: time.Now(),
// 	}
// }

// Showing the Creator or Constructor function, returning a pointer value instead. More efficient, prevent value from being copied.
// Not important for simple values like below
// func newUser(firstName, lastName, birthDate string) *user {
// 	return &user{
// 		firstName: firstName,
// 		lastName:  lastName,
// 		birthDate: birthDate,
// 		createdAt: time.Now(),
// 	}
// }

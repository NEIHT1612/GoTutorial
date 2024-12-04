package main

import (
	"fmt"
	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birth date (MM/DD/YYYY): ")

	var appUser user.User
	appUser = user.NewUser(userFirstName, userLastName, userBirthDate)
	var appAdmin user.Admin
	appAdmin = user.NewAdmin("thiendt@gmail.com", "123")

	outputUserDetails(appUser)
	outputAdminDetails(appAdmin)
}

func outputUserDetails(u user.User){
	fmt.Println(u.FirstName, u.LastName, u.BirthDate, u.CreateAt)
}

func outputAdminDetails(u user.Admin){
	fmt.Println(u.User.FirstName, u.User.LastName, u.Email, u.Password)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
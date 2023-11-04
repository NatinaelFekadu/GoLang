package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

type User struct {

	firstName string
	lastName string
	birthdate string
	createdDate time.Time
}

func main() {
	var newUser *User
	
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	newUser = NewUser(firstName,lastName,birthdate)
	
	newUser.outputDetails()
}

func NewUser(fName string, lName string, bDate string) *User{
	created := time.Now()

	newUser := User{
		firstName: fName,
		lastName: lName,
		birthdate: bDate,
		createdDate: created,
	}
	return &newUser
}

func(user *User) outputDetails(){
	fmt.Printf("My name is %v %v (born in %v)", user.firstName,user.lastName,user.birthdate)
}

func getUserData(promptText string) string{
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')

	cleanedInput := strings.TrimSpace(userInput)
	return cleanedInput
}


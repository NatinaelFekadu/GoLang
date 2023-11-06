package main

import (
	"bufio"
	"fmt"
	"os"
)

type Storer interface {
	Store(fileName string)
}

type Prompter interface {
	PromptForInput()
}

type PromptStorer interface{
	Storer
	Prompter
}

type userInputData struct {
	input string
}

type user struct{
	firstName string
	lastName string
	*userInputData
}

func newUser(firstName string, lastName string) *user{
	return &user{firstName: firstName,lastName: lastName,userInputData: &userInputData{},}
}

func newUserInputData() *userInputData {
	return &userInputData{}
}

func (usr *userInputData) PromptForInput() {
	fmt.Print("Your input please: ")
	reader := bufio.NewReader(os.Stdin)
	userInput,err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Fetching user input failed!")
		return
	}
	usr.input = userInput
}

func (usr *userInputData) Store(filename string){
	file,err := os.Create(filename)

	if err != nil {
		fmt.Println("Creating the file failed.")
		return
	}

	defer file.Close()
	file.WriteString(usr.input)
}

func main(){
	data := newUserInputData()

	nathy := newUser("Nathy","Fek")
	// data.PromptForInput()
	// data.Store("user1.txt")

	nathy.PromptForInput()
	nathy.Store("nathy.txt")
	handleUserInput(data)
}

func handleUserInput(container PromptStorer){
	fmt.Println("Ready to store your data!")
	container.PromptForInput()
	container.Store("user2.txt")
}

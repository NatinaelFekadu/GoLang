package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialAttackIsAvailable bool) string{
	
	for {
		playerChoice,_ := getPlayerInput()
		switch(playerChoice){
			case "1":
				return "ATTACK"
			case "2":
				return "HEAL"
			case "3":
				if specialAttackIsAvailable{
					return "SPECIAL_ATTACK"
				}
		}
		fmt.Println("Fetching the user input failed. Please try again.")
	}
}

func getPlayerInput() (userInput string, err error){
	fmt.Print("Your choice: ")

	userInput,err = reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)

	return
}
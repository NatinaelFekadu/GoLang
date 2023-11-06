package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

const (
	PlayerChoiceAttack = iota + 1
	PlayerChoiceHeal
	PlayerChoiceSpecialAttack
)

func GetPlayerChoice(specialAttackIsAvailable bool) string{
	
	for {
		playerChoice,_ := getPlayerInput()
		switch(playerChoice){
			case fmt.Sprint(PlayerChoiceAttack):
				return "ATTACK"
			case fmt.Sprint(PlayerChoiceHeal):
				return "HEAL"
			case fmt.Sprint(PlayerChoiceSpecialAttack):
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
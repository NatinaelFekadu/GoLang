package main

import (
	"github.com/none/project/actions"
	"github.com/none/project/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundData{}

func main() {
	startGame()

	winner := ""

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound % 3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)


	var playerAttackDamage int
	var playerHealValue int
	var monsterAttackDamage int

	switch(userChoice){
	case "ATTACK":
		playerAttackDamage = actions.AttackMonster(false)

	case "HEAL":
		playerHealValue = actions.HealPlayer()
		
	
	case "SPECIAL_ATTACK":
		playerAttackDamage = actions.AttackMonster(true)
	}

	monsterAttackDamage = actions.AttackPlayer()

	playerHealth,monsterHealth := actions.GetHealthAmounts()

	roundData := interaction.RoundData{
		Action: userChoice,
		PlayerHealth: playerHealth,
		MonsterHealth: monsterHealth,
		PlayerAttackDamage: playerAttackDamage,
		PlayerHealValue: playerHealValue,
		MonsterAttackDamage: monsterAttackDamage,
	}
	interaction.PrintRoundStatistics(&roundData)

	gameRounds = append(gameRounds, roundData)

	if playerHealth <= 0 {
		return "Monster"
	}else if monsterHealth <=0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
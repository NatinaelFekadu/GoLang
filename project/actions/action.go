package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currentMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) (damageValue int){
	minAttackValue := PLAYER_ATTACK_MIN_DAMAGE
	maxAttackValue := PLAYER_ATTACK_MAX_DAMAGE

	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN_DAMAGE
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_DAMAGE
	}
	damageValue = generateRandBetween(minAttackValue,maxAttackValue)
	currentMonsterHealth -= damageValue
	return

}

func HealPlayer() (int){

	healValue := generateRandBetween(PLAYER_HEAL_MIN_VALUE,PLAYER_HEAL_MAX_VALUE)

	healthDiff := PLAYER_HEALTH - currentPlayerHealth

	if healthDiff >= healValue {
		currentPlayerHealth += healValue
		return healValue
		
	}else{
		currentPlayerHealth = PLAYER_HEALTH
		return healthDiff
	}

}

func AttackPlayer() (damageValue int){
	damageValue = generateRandBetween(MONSTER_ATTACK_MIN_DAMAGE,MONSTER_ATTACK_MAX_DAMAGE)
	currentPlayerHealth -= damageValue
	return
}

func generateRandBetween(min int, max int) int{
	return randGenerator.Intn(max-min) + min
}

func GetHealthAmounts() (int, int){
	return currentPlayerHealth,currentMonsterHealth
}
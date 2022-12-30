package main

import (
	"fmt"
	"log"

	"github.com/treyarte/playground/helpers"
)

type Enemy struct {
	Health   float64
	Mana     string
	Strength int
	Damage   float64
}

func (attacker *Enemy) Attack(target *Enemy) (string, string) {
	if target.Health <= 0 {
		return "", "Target is dead"
	}

	totalDamage := float64(attacker.Strength/5.0) + attacker.Damage

	remainingHealth := target.Health - totalDamage

	var enemyStatus string

	if remainingHealth < -100 {
		 enemyStatus = "OVERKILL!!"
	}

	if remainingHealth < 0 && remainingHealth > -100 {
		enemyStatus = "You killed the enemy"
	}

	damageDealt := fmt.Sprintf("You did %v damage!", totalDamage)

	return damageDealt, enemyStatus
}

var max = 100

func CalculateValue(intChan chan int) {
	value := helpers.RandomNumber(max)

	intChan <- value
}

func main() {
	enemy1 := Enemy {
		Health: 980,
		Mana: "100",
		Strength: 63,
		Damage: 300,
	}

	enemy2 := Enemy {
		Health: 301,
		Mana: "100",
		Strength: 63,
		Damage: 300,
	}

	dmgDealt, enemyStatus := enemy1.Attack(&enemy2);

	fmt.Println(dmgDealt)
	fmt.Println(enemyStatus)

	intChan := make(chan int)
	defer close(intChan)
	
	go CalculateValue(intChan)

	log.Println(<- intChan)

}
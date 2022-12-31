package main

import (
	"encoding/json"
	"errors"
	"log"
)

type Enemy struct {
	Name string
	Health   float64
	Mana     float64
	Strength int
	AttackPower   float64
}

var validEnemy = Enemy{
	Name:        "Poring",
	Health:      150,
	Mana:        0,
	Strength:    1,
	AttackPower: 16,
}

var invalidEnemy = Enemy{
	Name:        "Blue Roda Frog",
	Health:      150,
	Mana:        0,
	Strength:    -1000,
	AttackPower: 200,
}

var Enemies = []Enemy {
	{
		Name: "Cell",
		Health: 900,
		Mana: 100,
		Strength: 78,
		AttackPower: 200,	
	},
	{
		Name: "Poring",
		Health: 150,
		Mana: 0,
		Strength: 1,
	},
	{
		Name: "Pupa",
		Health: 1000,
		Mana: 0,
		Strength: 0,
	},
}

var JsonEnemies string = `
	[
		{
			"Name": "Boomba",
			"Health": 600,
			"Mana": 0,
			"Strength": 26
		},
		{
			"Name": "Rappi",
			"Health": 480,
			"Mana": 20,
			"Strength": 15
		}
	]
`


func (attacker *Enemy) calculateDamage() (float64, error) {
	strBonusDmg := (attacker.Strength/20) * 4
	totalDamage := float64(attacker.Strength + strBonusDmg) + attacker.AttackPower

	if(totalDamage < 0) {
		return 0, errors.New("damage cannot be less than 0")
	}
	return totalDamage, nil;
}


func main() {

	var res []Enemy
	err := json.Unmarshal([]byte(JsonEnemies), &res)
	if err != nil {
		log.Fatal("Failed to convert from JSON")
	}

	log.Println(res)

	jsonStr, err := json.MarshalIndent(Enemies, "", "      ")
	if err != nil {
		log.Fatal("Failed to convert to JSON")
	}
	
	log.Panicln(string(jsonStr))

}
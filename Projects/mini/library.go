package main

import (
	"fmt"
)

const playerHP = 100
const playerPower = 10
const playerAttack = 10
const playerGold = 0
const playerLevel = 0

type Player struct {
	name   string
	hp     int
	power  int
	attack int
	level  int
	gold   int
}

type Rakchhas struct {
	name   string
	hp     int
	power  int
	attack int
	level  int
}

type MiniRakchhas struct {
	name   string
	hp     int
	power  int
	attack int
	level  int
}

// -------------------------------------- 1. rackchass method

/*

Basic rule i am thinking of now

1:
*/

func (Rakchhas) assign(level int) (Rakchhas, MiniRakchhas) {

}

// ----------------------------------- 2. players method

func (Player) assign(name string, hp int, power int, level int, attack int, gold int) Player {

	newPlayer := Player{name: name, hp: hp, power: power, level: level, attack: attack, gold: gold}

	return newPlayer
}

// although this isn't good design implementing game question logic in players method is wrong but still

func (Player) getRaakchas(level int) (Rakchhas, MiniRakchhas) {

}

func main() int {

	var pName string
	fmt.Println("=====================================================================")
	fmt.Println("welcome to Rakch-sikari")

	fmt.Println("Enter player name : ")
	fmt.Scan(&pName)

	// go doesn't have char instead it have byte,
	//  byte (an alias for uint8),  rune (an alias for int32)

	var options byte
	fmt.Println("> 1 start new game ")
	fmt.Println("> 2 load game ")
	fmt.Println(">3  quit game")

	fmt.Scan(&options)

	if options == 1 {
		player := Player{pName, playerHP, playerPower, playerAttack, playerLevel, playerGold}
	} else if options == 3 {
		return 0
	}
	// will implement state save and load later

	fmt.Println("starting...")
	rackchass, minirackchhas := player.getRaakchas(player.level)

	fmt.Println("======================= Fight Begin ===========================")

	player.getRaakchas

}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNumber:=rand.Intn(6)+1
	fmt.Println(diceNumber)
	switch diceNumber{
	case 1:
		fmt.Println("Dice is 1,you can open")
	
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
	case 4:
		fmt.Println("you can move 4 spot")
	case 5:
		fmt.Println("you can move 5 =spot")
	case 6:
		fmt.Println("you can move 6 spot amd roll again")
	default:
		fmt.Println("dice not rolled properly roll again")

	}
	
}
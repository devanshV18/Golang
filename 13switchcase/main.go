package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("The Dice Game!")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You move 1 place")
	case 2:
		fmt.Println("You move 2 place")
	case 3:
		fmt.Println("You move 3 place")
		fallthrough
	case 4:
		fmt.Println("You move 4 place")
	case 5:
		fmt.Println("You move 5 place")
	case 6:
		fmt.Println("You move 6 place")
	default:
		fmt.Println("Invalid dice Number")
	}
	

}
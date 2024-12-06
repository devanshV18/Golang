package main

import "fmt"

func main() {
	fmt.Println("Welcome to golang array tutorial")

	//declaring
	var fruitList [4]string

	//adding values
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Orange"
	fruitList[3] = "Pineapple"

	//prints the array simply in square brackets
	fmt.Println("FruitList :", fruitList)

	//VVII -> the len function always prints 4 i.e. it prints the size allocated to it irrespective of how many spaces it actually takes
	fmt.Println("FruitList length:", len(fruitList))


	var menu = [3]string{"Chicken Tikka", "Chicken kebab", "Ahuna Mutton"}
	fmt.Println("Menu for today is", menu)
	fmt.Println("Menu for today is", len(menu))
}
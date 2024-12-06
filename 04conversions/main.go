package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizzeria")
	fmt.Println("Please rate the pizza ona  scale of 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	//conversions ->using strconv package -> two inputs - sting and bit size

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if(err != nil){
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 to your rating", numRating+1)
	}
}
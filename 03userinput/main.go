package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome User"
	fmt.Println(welcome)

	//creating a reader using go packages
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our Pizza")

	//comma ok syntax (alternate to try catch) to handle input
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,",input)
	fmt.Printf("Type of this rating is %T", input)
}
package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 25
	var result string

	if loginCount < 10{
		result = "Regular User"
	} else if loginCount>10 {
		result = "Sus"
	} else{
		result = "Exactly 10 "
	}

	fmt.Println(result)


	if num := 11; num < 10 {
		fmt.Println(num)
	}else{
		fmt.Println("Num is not less than 10")
	}
}
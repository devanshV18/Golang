package main

import "fmt"

//main function is the entry point of every golang file
//we cant define our custom function inside main fxn just like cpp
func main() {
	greet()
	greetTwo()
	fmt.Println("Functions in Golang")

	result := adder(3,5)
	fmt.Println("Result is", result)

	proResult, myMessage := proAdder(2,5,7,8,2,4)
	fmt.Printf("Pro Result is %v and myMessage is %v", proResult, myMessage)
}

func greetTwo()  {
	fmt.Println("Hello from Golang")
}

func adder(valOne int, valTwo int) int{
	return valOne + valTwo
}

func greet(){
	fmt.Println("namastey from Golang")
}

func proAdder(values ...int) (int, string){
	total := 0
	for _, val := range values{
		total+=val
	}
	return total, "Hi from pro adder"
}


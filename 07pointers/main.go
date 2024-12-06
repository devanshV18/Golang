package main

import "fmt"

func main() {
	fmt.Println("Pointers Tutorials")

	//initialising an int pointer named as ptr
	var ptr *int 
	//NOTE -> san uninitialised pointer has a value nil
	fmt.Println("Value of a Pointer is ", ptr)



	//creating a ptr for a memory allocated using walrus op
	myNumber := 23
	var ptr2 = &myNumber

	//printing hexadecimal address
	fmt.Println("Value stored by pointer ", ptr2)

	//dereferencing ->same as cpp
	fmt.Println("Value stored at the memory address stored by pointer ", *ptr2)


	//updating value using ptr

	*ptr2 = *ptr2 + 10

	fmt.Println("New value incremented by 10 to", myNumber)
}
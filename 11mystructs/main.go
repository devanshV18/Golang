package main

import "fmt"

//there is no inheritance in golang hence no suoer, child or parent.

func main() {
	fmt.Println("Structs in Golang")

	devansh := User{"Devansh", "devanshverma024@gmail.com", true, 21}
	fmt.Println(devansh)

	fmt.Printf("Devansh's details are: %+v\n", devansh)
	
	fmt.Printf("Name is %v and Email is %v\n", devansh.Name, devansh.Email)
}

//first capital letter make them public
type User struct{
	Name string
	Email string
	Status bool
	Age int
}
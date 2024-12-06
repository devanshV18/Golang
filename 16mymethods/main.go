package main

import "fmt"

//there is no inheritance in golang hence no suoer, child or parent.

//functions inside structs is methods

func main() {
	fmt.Println("Structs in Golang")

	devansh := User{"Devansh", "devanshverma024@gmail.com", true, 21}
	fmt.Println(devansh)

	fmt.Printf("Devansh's details are: %+v\n", devansh)
	
	fmt.Printf("Name is %v and Email is %v\n", devansh.Name, devansh.Email)

	devansh.GetStatus()

	//tp chcek if NewMail changes original data or not
	devansh.NewEmail()
	fmt.Printf("Name is %v and Email is %v\n", devansh.Name, devansh.Email)

}

//first capital letter make them public


type User struct{
	Name string
	Email string
	Status bool
	Age int
}

//this method can be above or below struct but out of struct
func (u User) GetStatus() {
	fmt.Println("is user active:", u.Status)
}

//VVVIII
//does this below function change the actual value completely or just prints a copy?
//NO it does not, hence u User being passed makes the copy of the object and pass it to the function whenever we use it with any object.
func (u User) NewEmail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user set to", u.Email)
}
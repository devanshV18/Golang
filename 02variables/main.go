package main

import "fmt"

// The walrus operator (:=) only works inside a function

//hence below line throws an error ->
//jwtToken := 857857hgeuth

//vvvvii ->

//unchangable values -> constant

//the below variable name starts with a capital letter (L) which means this is a public variable -> In golang this means that the LoginToken variable is accessible to any file in this folder or to anyone throughout this program.

const LoginToken = "gibberishwu857gibberish"

func main() {
	var username string = "Devansh"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("variable is of type: %T \n", smallValue)

	var small int = 256
	fmt.Println(small)
	fmt.Printf("variable is of type: %T \n", small)

	var smallFloat float32 = 255.54475567567
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type %T \n", smallFloat)

	var preciseFloat float64 = 255.54475567567
	fmt.Println(preciseFloat)
	fmt.Printf("variable is of type %T \n", preciseFloat)

	//default values and aliases
	var anotherVariable int //default on a int is 0 -> no garbage
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type %T \n", anotherVariable)

	// implicit type -> no explicit mention but lexer automatically decides it in case of implicit types
	var webiste = "google.com"
	fmt.Println(webiste)

	//no var style -> by using := symbol (works only inside a function scope)
	// := is called walrus operator.
	grade := "Grade1"
	fmt.Println(grade)


	//LoginToken is accessed below as it is a public variableṇ
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type %T \n", LoginToken)
}
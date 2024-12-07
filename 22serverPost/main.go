package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//in this file, we are going to use a basic node js server with a few routes and handle their reponse and http methods using golang while keeping the backend server up and running.

func main() {
	fmt.Println("Handling server responses from a node js server upp and runniing -> POST")
	performPostJsonRequest()
}
func performPostJsonRequest(){
	const myurl = "http://localhost:8000/post"

	//fake json payload

	// string is the package which has a NewReader function
	//New Reader function is the function of strings package where we can use backticks and create a new reading format for our data. In the below case, we used {} inside `` to create a json view data.

	requestBody := strings.NewReader(`
		{
			"coursename":"Lets go with Golang",
			"price": 8999,
			"platform": "learncodeonline.in"
		}
	`)

	//we need url, contenttype, paylaod
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil{
		panic(err)
	}
	defer response.Body.Close()

	//no error lets read response
	content, _ := ioutil.ReadAll(response.Body)

	// fmt.Println(content)//this prints the bytes form of data
	
	//this prints the actual form of data
	fmt.Println(string(content)) 
}
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//in this file, we are going to use a basic node js server with a few routes and handle their reponse and http methods using golang while keeping the backend server up and running.

func main() {
	fmt.Println("Handling server responses from a bnode js server upp and runniing")


	PerformGetRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Satsus code", response.StatusCode)
	fmt.Println("Content Length is", response.ContentLength)


	//SCENARIO1 using usual method way1
	//note reading content the below line is inevitable, we have to do this everytime without any doubt
	//NOTE -> content stores the read file in bytes
	// content,_ := ioutil.ReadAll(response.Body)

	//converting content(byte format) to string
	// fmt.Println(string(content))

	//SCENARIO2 -> using sring builder way2
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Bytecount is", byteCount)

	fmt.Println(responseString.String())

}
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//VVVVIII ->
//in golang we handle web req using net/http library, this library has a response object that is a struct containing key value pair such as status, status codes, Headers, Body, contentLength etc so that we can access data more conviniently.

//the url wont work obviously
const url = "https://lco.dev"

func main() {
    fmt.Println("Lco web requuest tutorial")

    //Get() only takes one argument
    response, err := http.Get(url)

    if err != nil {
        panic(err)
    }

    fmt.Printf("Response type is %T\n", response)

    //the Body as the name indicates has a capital start letter which means Body is a public variable and it has to be cllosed manually.
    defer response.Body.Close() 

    //to read 
    databytes, err := ioutil.ReadAll(response.Body)

    if err != nil {
        panic(err)
    }
    //if no error lets change our content into string
    content := string(databytes)

    //printing content
    fmt.Println((content))

}
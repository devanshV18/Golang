package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	performPostFormRequest()
}

func performPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formData
	//data is picked using url packages

	data := url.Values{}

	data.Add("firstname", "Devansh")
	data.Add("lastname", "Verma")
	data.Add("email", "dv@google.dev")

	//post form is the function used to handle form data -> 
	//vvviii -> PostForm issues a POST to the specified URL, with data's keys and values URL-encoded as the request body.
	// The Content-Type header is set to application/x-www-form-urlencoded. To set other headers, use [NewRequest] and DefaultClient.Do.

	//url, data
	response, err := http.PostForm(myurl, data)

	if err!= nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(response.Body)

	defer response.Body.Close()

	fmt.Println(string(content))
}
package main

import (
	"fmt"
	"net/url"
)


const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=uriuw765"

func main() {
	fmt.Println("Handling URL in Golang")
	fmt.Println(myurl)


	//practically we will parse the url
	result, _ := url.Parse(myurl)

	//whenever we parse a url and store it in a variable, here result, result has alot of information to explore.

	fmt.Println(result.Scheme) //prints https
	fmt.Println(result.Host) //lco.dev:3000
	fmt.Println(result.Path) // /learn
	fmt.Println(result.Port()) // 3000
	fmt.Println(result.RawQuery) //coursename=reactjs&paymentid=uriuw765

	//NOTE RawQuery is equivalent to params in js


	//VVVVIIII ->
	//storing params, Raw query in a variable qparams => stores the params in a key value pair format
	qparams := result.Query()
	fmt.Printf("The type of qparams is %T\n", qparams) //url.Values -> key value pairs


	//using key value pair to extract info from qparams.
	//VVVIII
	fmt.Println(qparams["coursename"]) //prints the value of this parameter key. -> reactjs
	fmt.Println(qparams["paymentid"]) //prints the value of this parameter key. -> reactjs


	//LOOPING THROUGH QPARAMS:-
	for key, val := range qparams{
		fmt.Printf("Key is %v and Value is %v\n", key, val)
	}


	//working the exact other way round -> building a url from chunks

	//the exact syntax to pass chunks in key value pairs and construct our url
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "sfcl.mar",
		Path: "/race",
		RawPath: "driver=charles",
	}

	// storing our contructed url in a variable using String()
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
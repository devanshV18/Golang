package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //using coursename alias for Name in Json format
	Price    int
	Platform string `json:"website"`
	Password string `json:"-"` //removes from resp
	Tags     []string
}

func main() {
	fmt.Println("Handling/Consuming Json")
	DecodeJson()
}

//NOTE -> Whenever data is received from web it is also=ways in byte format which we convert to string but this time we will consume it in JSON format as it is coming in json form

func DecodeJson() {
	//[]byte -> our json is array of obj and since we receive it in bytes inititally.

	jsonDataFromWeb := []byte(`
		{
            "coursename": "ReactJs Bootcamp",
            "Price": 299,
            "website": "learncodeonline.in",
            "Tags": ["webdev","js"]
        }
	`)

	//this is the variable / interface that holds our json unmarshalled data:- METHOD 1
	var lcoCourse course 

	//checking if the json is valid or not
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Valid Json")
		//unmarshal takes two arg -> data and an interface where interface is the reference of a variable of type of struct
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)

		//%#v to print json interface variables
		fmt.Printf("%#v", lcoCourse)
	}
	
	fmt.Println()
	//METHOD2
	//in some cases we want to add values as key value

	var myOnlineData map[string]interface{}

	//we can use the error handling stuff again but just doing it right away

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	//prints in form of keyvalue -> since its a map, the order might be jumbled vastly
	fmt.Printf("%#v", myOnlineData)

	//looing our map

	for key, value := range myOnlineData{
		fmt.Printf("Key is %v  and Value is %v ", key, value)
	}
}
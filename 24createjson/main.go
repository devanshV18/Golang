package main

import (
	"encoding/json"
	"fmt"
)

//created a struct for creating our json
type course struct{
	Name string `json:"coursename"` //using coursename alias for Name in Json format
	Price int
	Platform string `json:"website"`
	Password string `json:"-"` //removes from resp
	Tags []string
}

func main() {
	fmt.Println("Welcome to Json video")
	EncodeJson()
}

//json encoding -> in this step we convert our available data such as slices, key value pairs etc into json

func EncodeJson()  {
	//vvvi
	// lcoCourses is our data -> which is a slice containing values of type course ->the struct we created, as mentioned below

	lcoCourses := []course{
		{"ReactJs Bootcamp", 299, "learncodeonline.in", "abc123", []string{"webdev", "js"}},

		{"MERN Bootcamp", 299, "learncodeonline.in", "def123", []string{"fullstack", "ts"}},

		{"ML Bootcamp", 299, "learncodeonline.in", "mnoc123", []string{"AI/ML", "py"}},
	}

	//package this data as json data
	//will be done using json library and Marshal function -> Marshal returns the JSON encoding of v.

	// Marshal traverses the value v recursively. If an encountered value implements [Marshaler] and is not a nil pointer, Marshal calls [Marshaler.MarshalJSON] to produce JSON. If no [Marshaler.MarshalJSON] method is present but the value implements encoding.TextMarshaler instead, Marshal calls encoding.TextMarshaler.MarshalText and encodes the result as a JSON string.

	//NOTE -> some times it might not be possible to convert our passed data to json, so we must handle err using comma syntax.

	//conversion is abit hard to read with the below line
	// finalJson, err := json.Marshal(lcoCourses) 

	//better way 2 ->using MarshalIndent
	// args -> data, prefix(no need unnecessary), indent
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") 
	

	if err != nil{
		panic(err)
	}

	//no error

	//prints an array of objects similar as json
	fmt.Printf("%s\n", finalJson) //a bit hard to read
}
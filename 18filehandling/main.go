package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file Handling in Golang")

	content := "This needs to go in a file - Learncoding.in"

	// file has a path if file is created successfully and error has the error if something goes wrong
	file, err := os.Create("./mylcogofile.txt")

	//handling err
	checkNilErr(err)
	//if no error lets write and get length
	length,err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("The length is", length)
	defer file.Close()

	readFile("./mylcogofile.txt")
}



//lets read a file

func readFile(filename string)  {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("text data inside the file is \n", databyte)
}


func checkNilErr(err error)  {
	if err != nil{
		panic(err)
	}
}
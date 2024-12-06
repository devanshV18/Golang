package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to out time package tutorial in Golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	//format works wrt to this date in golanfg
	fmt.Println(presentTime.Format("01-02-2006"))//to get the date

	fmt.Println(presentTime.Format("01-02-2006 Monday")) //to get date and day

	//to get current ticking time
	fmt.Println(presentTime.Format("01-02-2006 15:04:05")) 


	//to get everything
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) 


	//creating time -> the arguments need to be like this only
	createdDate := time.Date(2020, time.December, 10, 11, 45, 54,0, time.UTC)
	fmt.Println(createdDate)

	//using formatting to print desired part of our created date
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
package main

import (
	"fmt"
	"sort"
)

//slices are arrays + features + some abstraction

//its important to get strong hold of slices to work with DB deffectivey

func main()  {
	fmt.Println("Welcome to slice Tutorials")

	//1. we added 3 values without explicitly declaring size
	var fruitList = []string{"Apple", "Peach", "tomato"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	// hence in slices we can add as many values as we want to, it dynamically expands

	//adding value
	//3. we added two values to the fruitList slice but we need to use this var again, else it throws an error
	fruitList = append(fruitList, "Mango", "Guava")
	fmt.Println(fruitList)

	//4. the below operation removed first value from the slice
	//starts at index 2 goes till end
	// fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)

	//5. VVVII
	//note -> the below operation occurs on the result of first opearation carried above, not on original slice


	//5. another example [remove,stop] non inclusiv ranges
	//start from 2nd element (1 skipped) go till before 3rd
	//first arg is the posn and second is index
	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)


	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	//using amke keywowrd 
	//args -> type of data and size
	highScores := make([]int, 4) //this is declared slice

	highScores[0] = 154
	highScores[1] = 124
	highScores[2] = 112
	highScores[3] = 101

	fmt.Println(highScores)

	//asper the declaration of this slice, it seems that if we use append functiona nd add a few more highscores, its gonna throw an error but it is not the case because make gives us a slice definately but when we use append the memory managemnt reallocated amemory and accomodate any incoming value

	highScores = append(highScores, 166, 178, 183)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted((highScores)))
	sort.Ints(highScores)
	fmt.Println(highScores)

	//to remove a value from slices based on index

	var courses = []string{"React", "js", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
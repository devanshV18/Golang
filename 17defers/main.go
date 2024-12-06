package main

import "fmt"

func main() {
	//defer keywords basically puts the line of execution to the very last position before the main function closing braces

	// defer fmt.Println("Hello")
	// fmt.Println("Understanding defers")

	//scenario 2 -> defers are palced in reversed order(LIFO)
	//world defer is at last, just above comes one, just above comes two, hello prints, two exec, one exec, world exec(LIFO)
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("hello")


	//scenario3 -> a defer function
	myDefer()

	//VVVVVIIIIII -> 3 defers go to end, hello is printend , when myDefer is executed ->
	// i = 0 defer, i= 1, defer, i= 2 defer, i =3 defer, i=4 defer.

	// hence defer order: -
	//4
	//3
	//2
	//1
	//0
	//two
	//one
	//world

	// once hello is printed the defer stack start execution in lifo order i.e. hello -> 4 -> 3 -> 2 -> 1 -> 0 -> two -> one -> world
}

func myDefer(){
	for i := 0; i<5; i++ {
		defer fmt.Println(i)
	}
}
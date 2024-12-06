package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	days := []string{"Saturday", "Friday", "Monday", "Wednesday"}

	fmt.Println(days)

	//1. loop1
	// for d:=0; d<len(days); d++{
	// 	fmt.Println(days[d])
	// }


    //2. loop2
	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	//3. loop3 -> we can drop either of index, day using _
	for index, day := range days{
		fmt.Printf("Index is %v and value is%v\n", index, day)
	}

	rougeValue := 1

	//4. while loop in disguise of for loops
	for rougeValue < 10 {
		if rougeValue == 5 {
			// break
			rougeValue++
			continue
		}

		if rougeValue == 7{
			goto lco
		}
		 fmt.Println("Value is: ", rougeValue)
		 rougeValue++
	}

	fmt.Println("LOOP KI MKC")

	lco:
		fmt.Println("jumping at amazon.in")


}
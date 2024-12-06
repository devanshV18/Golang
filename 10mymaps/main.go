package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["CPP"] = "c++"
	languages["RB"] = "ruby"
	languages["PY"] = "python"

	//all key value pairs get automatically sorted and they aren't separated using comma
	fmt.Println("List of all Programming languages", languages)

	fmt.Println("JS Shorts for", languages["JS"])

	//deleting entry from map

	delete(languages, "RB")
	fmt.Println(languages)


	//loops in golang for maps
	for key, value := range languages {
		fmt.Printf("For key %v, Value is %v\n", key, value)
	}

	//demonstrating comma ok syntax here
	//when we dont need key we can assign _ to it
	for _, value := range languages {
		fmt.Printf("Value is %v\n",value)
	}
}
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Mod in Golang")
	greet()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET") //the route server only GET req on / route if we add the methods part like this

	//the below line runs a server with our port and router as input
	//NOTE -> we can handle error using comma ok syntac or using log fatal below
	log.Fatal(http.ListenAndServe(":4000", r)) //handles the error  as well as runs the server
}

func greet()  {
	fmt.Println("hello MOD Users")
}

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to Golang tutorial</h1>"))
}
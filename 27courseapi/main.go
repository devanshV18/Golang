package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//action paln -> slice will the dummy db

//Model for course
type Course struct{
	CourseId string `json:"cousreid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author //type of Author is Author Pointer
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}


//fake db
var courses []Course

//middlewares or helpers
func (c *Course) IsEmpty() bool  {
	return c.CourseName == "" //return true if coursename is empty
}


//controllers -> separate files

//1. serve home router
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to Courses API</h1>"))
}

//2. all courses
func getAllCourses(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Get all courses")
	//setting headers for this controllers -> below line means this controller json content
	w.Header().Set("Content-Type", "application/json")

	//throwing the entire slices
	//New encoder works with a  writer, w in our case and we use encode fxn to encode the courses which means the argument passed in Encode() will be treated as json
	json.NewEncoder(w).Encode(courses)
}

//3. get one course
//NOTE -> ENCODING IS DONE TO CREATE SEND A JSON REPSONSE
func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from req ->url package can also be used
	params := mux.Vars(r) //this line extracts all variables from the url to params variable
	
	//loop thru courses and pick the matching course
	for _, course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course) //futher destructuring can be done to send more specific data

		}
	}

	//no course found case
	json.NewEncoder(w).Encode("No course found with that id")

}

//4.create one course
//NOTE -> Decoding is done to accept a json req
func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create one course")
	w.Header().Set("Content-type", "application/json") //setting header

	//checking if req has a json or is empty
	if r.Body == nil{
		json.NewEncoder(w).Encode("Please send json data")
	}

	//an empty json {}
	//handle json -> just as we learnt doing in consume json tutorial
	var course Course

	//decoding to accept and process a json req
	_ = json.NewDecoder(r.Body).Decode(&course) 

	if course.IsEmpty(){
		json.NewEncoder(w).Encode("Empty json data") //sending resp
		return
	}

	//to generate a unique course id
	//append this new course to the slice

	rand.Seed(time.Now().UnixNano())
	//rand.Intn(100) gives us a random int for id, we use strconv.itoa fxn to convert int to string in go and save it in the CourseId field of our course.
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course) //final resp, we can also send messages etc (auto return)
}

//5
func updateOneCourse(w http.ResponseWriter, r *http.Request){
	//loop and pic the course to be updated
	fmt.Println("Create one Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop and remove , add operation again from id destructured above

	for index, course := range courses{
		if course.CourseId == params["id"]{
			//deleting from slices
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course) //setting up course interface with the decoded json values from request
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("The targeted course does not exist")
}

//6 delete one course
func deleteOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Delete one Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted the Targeted course")
		}
	}

	json.NewEncoder(w).Encode("The targeted course dows not exist")
}


func main()  {
	fmt.Println("Courses API")
	r := mux.NewRouter()

	//seeding 
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 345, Author: &Author{Fullname: "Devansh Verma", Website: "dv.dev"}})

	courses = append(courses, Course{CourseId: "4", CourseName: "NodeJS", CoursePrice: 335, Author: &Author{Fullname: "Reena Verma", Website: "rv.dev"}})

	courses = append(courses, Course{CourseId: "6", CourseName: "ExpressJS", CoursePrice: 325, Author: &Author{Fullname: "Mani Bhushan Lall", Website: "mbl.dev"}})

	courses = append(courses, Course{CourseId: "8", CourseName: "NextJs", CoursePrice: 315, Author: &Author{Fullname: "Randeep Hooda", Website: "rh.dev"}})

	//router
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")


	//listen
	log.Fatal(http.ListenAndServe(":4000", r))

}
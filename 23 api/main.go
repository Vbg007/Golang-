package main
// model for courses-file

import path {
	"fmt"
	"net/http"
	"encoding/json"
	"math/rand"
	"time"
	"strconv"
}
type Course struct{
	CourseID string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"courseprice"`
	Author *Author `json:"author"`

}
type Author struct{
   Fullname string`json:"fullname"`
   Website string `json:"website"`

}
//mock db
var courses []Course
// middle ware, helper
func(c *Course) IsEmpty()bool{
   return c.CourseId == "" &&c.CourseName ==""
}
func main (){
	

}
//controllers
// serve home route
func servehome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1> Welcome to api by learn code online </h1>"))
}
func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)

}
func getOneCourse(w http.ResponseWriter, r *http.Request){ 
	fmt.Println("Get one course")
	w.Header().Set("Content-Type","application/json")
    // grab id from request
	params := mux.Vars(r)
	// loop through courses and find with id
	
	for _, course := range courses{
		if course.CourseID == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("coursenot found")
	return

}
func CreateOneCourse (w.http.ResponseWriter, r *http.Request){
         fmt.Println("Create one course")
		 w.Header().Set("Content-type","application/JSON")
		 if r.body== nil {
			json.NewEncoder(w).Encode("new Course")
		 }
         // for empty json {}
		 var course Course
		 _= json.NewDecoder(r.body).Decode(&course)
		 if course.IsEmpty(){
			json.NewEncoder(w).Encode("No data found")
			return
		 }
        // generate unique id, string
		// append coirse to Courses
		rand.Seed(time.Now().UnixNano())
		course.CourseID = strconv.Itoa(rand.Intn(100))
		courses=append (courses,course)
		json.NewEncoder(w).Encode(course)
		return
	}  
 func UpdateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update one course")
	w.Header().Set("Content-type","application/JSON")
	// grab id from request
	params := mux.Vars(r)
	for index ,course := range courses{
		if course.CourseID == params["id"]{
            courses=append(courses[:index],courses[index+1]...)
			_=json.NewDecoder(r.body).Decode(&course)
			course.CourseID=params["id"]
			courses = append(courses,course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
 }
 func deleteOneCourse(w.httpwriter,r *http.request){
	fmt.Println("Delete one course")
	w.Header().Set("Content-type","application/JSOn")
	 params :=mux.Vars( r)
	 for index,course := range courses{
		if course.CourseID == params["id"]{
			courses= append(courses[:index],courses[index+1:]...)
			break
		}
	 }
	
 }
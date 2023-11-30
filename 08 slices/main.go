package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("slices")
	var mylist =[]string{"Apple","Tomato","Banana"}
	fmt.Println(len(mylist))
	mylist=append(mylist,"Mongo","Raspberrry")
	fmt.Println(mylist)
	//slicing is same as python
	mylist=append(mylist[:4])
	fmt.Println(mylist)
	highscores:=make([]int,4)
	highscores[0]=26
	highscores[1]=46
	highscores[2]=17
	highscores[3]=2
	highscores= append(highscores,33,1,80)
	
	fmt.Println((highscores))
	sort.Ints(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))
	fmt.Println(highscores)
	courses:= []string{"react","java","python"}
    fmt.Println(courses)
	index:=2
	courses =append(courses[:index],courses[index+1:]...)
    fmt.Println(courses)
	 

}
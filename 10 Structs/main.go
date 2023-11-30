package main

import "fmt"

func main() {
	fmt.Println("Structs")
	// no inheritance in golang:no super or parent
 vbg:=User{"V B Gnanavel","gnanavelvb.20eee@kongu.edu",46,"Intern"}
 fmt.Println(vbg)
 fmt.Printf("%v",vbg.Id)
}  

type User struct{
	Name string
	Emain string
	Id int
	Position string
}
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Structs")
	// no inheritance in golang:no super or parent
 vbg:=User{"V B Gnanavel",26,"Intern","gnana@gmail.com"}
 fmt.Println(vbg)
 vbg.GetStatus()
 vbg.NewMail()
 fmt.Println(vbg)
}  

type User struct{
	Name string
	Id int
	Position string
	email string
}

func (b User) GetStatus(){
    fmt.Println("IS user active: ",b.Id)
}

func (u *User) NewMail(){
	reader:= bufio.NewReader(os.Stdin)
     u.email,_ = reader.ReadString('\n')

}
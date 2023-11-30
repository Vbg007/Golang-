package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("files")
	content:="I am writing to you my Happiness "
	file,err:= os.Create("./mylovefile.txt")
	checkNilErr(err)
	length,err:=io.WriteString(file,content)
	checkNilErr(err)
	fmt.Println(length)
	file.Close()
	readFile("mylovefile.txt")
	
}
func readFile(filname string){
	databyte,err:=os.ReadFile(filname)
    checkNilErr(err)
	fmt.Println(databyte)
}
func checkNilErr(err error){

	if err!=nil{
		panic(err)//shutdown the program while error and shos error
	}

}
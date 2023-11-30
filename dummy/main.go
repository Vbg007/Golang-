package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
 fmt.Println("Hello, World!")
 reader := bufio.NewReader(os.Stdin)
 fmt.Print("Enter text: ")
 input,err :=reader.ReadString('\n')
 if err != nil{
   	fmt.Println("An error occured while reading input. Please try again",err)
 	return
 }
 fmt.Println("Input was: ",input)

}
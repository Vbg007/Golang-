package main

import "fmt"

func main() {
	defer fmt.Println("thalaivarey")
	fmt.Println("hello");
	myDefer()
}
func myDefer(){
	for i:=0;i<5;i++{
		defer fmt.Println(i)

}}
package main

import "fmt"

func main() {
	fmt.Println("pointers")
	// var ptr *int
	// fmt.Println(ptr)

	text := "vbg"
	var ptr = &text
	fmt.Println(*ptr)// pointer is same as c
}

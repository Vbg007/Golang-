package main

import "fmt"

func main() {
	var fruitlist [4]string
	fruitlist[0]="apple"
	fruitlist[1]="banana"
	fruitlist[2]="carrot"
	fmt.Println(fruitlist[2])
	fmt.Println(len(fruitlist))// only shows the declared size not used size
	var veglist[4] string{"tomato","cabbage"}
}
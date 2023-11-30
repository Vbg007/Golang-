package main

import "fmt"

func main() {
	fmt.Println("maps")
	languages:=make(map[int64]string)

	languages[21]="javascript"
	languages[22]="Ruby"
	languages[23]="python"
	fmt.Println(languages)
	fmt.Println(languages[21])
	delete(languages,22)
	fmt.Println(languages)
	for key,value := range languages{
		fmt.Printf("key is %v, value is %v\n",key,value)
	}

}
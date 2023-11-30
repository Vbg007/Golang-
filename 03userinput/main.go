package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	userinput := "Vanakkam da mapla"
	fmt.Println(userinput)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age")
	//comma ok or comma error syntax
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	fmt.Printf("%T",input);
	num := 5
	num2 := 21
	var num3 int = num + num2
	fmt.Println(num3)

}

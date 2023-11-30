package main

import "fmt"
func greeter(){
	fmt.Println("Vanakkam da Mapla")
}
func adder(x int, y int) int{
	return x+y
}
func main() {
	fmt.Println("vbg");
	greeter()
	result:=adder(5,4)
	fmt.Println(result)
	proResult,Mymessage:=proAdder(26,27)
	fmt.Println(proResult)
	fmt.Println(Mymessage)

}
func proAdder(values...int)(int,string){
	total:=0
	for _,val:=range values{
		total+=val
	}
	str:="Leo ooh"
	return total,str
}
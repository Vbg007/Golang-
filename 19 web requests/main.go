package main

import (
	"fmt"
	"io"
	"net/http"

)

const url ="https://www.yahoo.com/"
func main() {
	fmt.Println("LCO web requests")
	response,err := http.Get(url)
	checkerr(err);
	fmt.Printf("%T",response)
	defer response.Body.Close()
	databytes,err := io.ReadAll(response.Body)
	checkerr(err);
	content := string(databytes)
	fmt.Println(content)

}
func checkerr(err error){
	if err!=nil{
		panic(err)
	}
}

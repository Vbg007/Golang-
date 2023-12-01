package main

import (
	"fmt"
	"io"
	"net/http"
	"github.com/gorilla/mux"
)

const url ="https://www.yahoo.com/"
func main() {
	fmt.Println("LCO web requests")
	response,err := http.Get(url)
	checkerr(err);
	fmt.Printf("%T",response)
	defer response.Body.Close()
	databytes,err := io.ReadAll(response.Body)
    if err!=nil{
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

}


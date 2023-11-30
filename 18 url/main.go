package main

import (
	"fmt"
	"net/url"
)
const myurl string="https://lco.dev:3000/learn?coursename=reactjs&paymentis=ghbj456ghb"

func main() {
	fmt.Println("Handling URL")
	//parsing
	result,err := url.Parse(myurl)
	checkerr(err)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)
	qparams := result.Query()
	fmt.Printf("%T",qparams)
	for _,val := range qparams{

		fmt.Println("Param is: ",val)
	}
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "lco.dev",
	    Path:   "/tutcss",
		RawPath:"user=hitesh",

	}
	anotherURl := partsOfUrl.String()
	fmt.Println(anotherURl)
   
}
func checkerr(err error){
	if err!=nil{
		panic(err)
	}
}
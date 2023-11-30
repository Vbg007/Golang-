package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web verb video")
//	PerformGetRequest()
	// PerformPostJSONRequest()
    performPostformRequest()
}


func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
		response, err := http.Get(myurl)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		fmt.Println("Status code: ", response.StatusCode)
		fmt.Println("Content length is:", response.ContentLength)

		var responseString strings.Builder
		content, _ := io.ReadAll(response.Body)
		bytecount, _  := responseString.Write(content)
		fmt.Println(bytecount)
		fmt.Println(responseString.String())

	 
}
func PerformPostJSONRequest() {
	const myurl = "http://localhost:8000/post"
	
	requestBody := strings.NewReader(`
	{
		"Name":"V B Gnanavel",
		"price": 0,
		"platform":"learncode0nline.in"
	}
`)
	response,err := http.Post(myurl,"application/json",requestBody)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	content, _ :=io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func performPostformRequest(){
  const myurl="http://localhost:8000/postform"
  data :=url.Values{}
  data.Add("name","vbgnanavel")
  data.Add("lastname","balamurugan")
  data.Add("email","vbg123@go.dev")
  response,err := http.PostForm(myurl,data)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	content, _ :=io.ReadAll(response.Body)
	fmt.Println(string(content))


}

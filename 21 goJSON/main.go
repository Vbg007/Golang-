package main

import (
	"encoding/json"
	"fmt"
)
type course struct{
	Name     string `json:"coursename"`//aliasing the key name in jso
	Price    int
    Platform string `json:"-"`// to hide the key in json
	Password string `json:"-"`
	Tags     []string  `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to json video")
	//EncodeJson()
	DecodeJson()
}
func EncodeJson(){
	courses := []course{
		{"Reactjsbootcamp",299,"iaminto","abc123",[]string{"web-dev","js"}},
		{"Golang",299,"iaminto","abc123",[]string{"go-lang","js"}},
		{"Angular",299,"iaminto","abc123",[]string{"web-dev","angular-js"}},
		{"Angular",299,"iaminto","abc123",[]string{}},
		
	}

	finalJson,err := json.MarshalIndent(courses,"","\t")
    checkerr(err)
	fmt.Println(string(finalJson))

}
func checkerr(err error){
	if err !=nil{
		panic(err)
	}
}
func DecodeJson() {
	jsonFromWeb := []byte(`
   {
			"coursename": "Reactjsbootcamp",
			"Price": 299,
			"tags": [
					"web-dev",
					"js"
			]
   } 
	`)
	var vbgcourse course
	checkvalid :=json.Valid(jsonFromWeb)
	if checkvalid{
		fmt.Println("valid json")
		json.Unmarshal(jsonFromWeb, &vbgcourse)//if use var name it makes acopy so store on actuual var 
	
	} else {
		fmt.Println("jsonis not valid")
	}
	var myOnlinedata map[string]interface{}
	json. Unmarshal(jsonFromWeb,&myOnlinedata)
	fmt.Printf("%#v\n",myOnlinedata)
	for key,value := range myOnlinedata{
		fmt.Printf("key is %v value is %v and type is %T\n",key,value,value)

	}
}
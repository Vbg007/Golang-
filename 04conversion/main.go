package main

import ( 
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to petrus\n Enter your Id no")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("input:", input)
	id, err := strconv.ParseInt(strings.TrimSpace(input), 10,64)
	fmt.Println("id: ", id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added init to your Id: ", id+1)
	}
}

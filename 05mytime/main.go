package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime.Format("date: 01-02-2006\nDay:Monday\nTime:15:04:05"))
	// custom created date
	createdDate := time.Date(2023, time.September, 02, 8, 46, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("A special day in my memory\nDate: 02-01-2006\nDay: Monday\nTime: 15:04:05"))

	
}

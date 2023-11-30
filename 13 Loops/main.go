package main

import "fmt"


func main() {
	days:= []string{"sunday","Monday","Saturday","Valentines day","her Birthday"}
	fmt.Println(days)

	// for d:=0;d<len(days);d++{
	// 	fmt.Println("Today is "+days[d])
	// }
	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	for index,day:=range days{
		fmt.Printf("Index is %v nd day is %v\n",index,day)
	}
	rogue:=1
	for rogue <10{
		if rogue==3{
			goto lcu
		}
		if rogue ==5{
			rogue++
		    break

		}
		fmt.Println(rogue)
		rogue++
	}
lcu:
   fmt.Println("Jumping at lokesh cinimatic universe")

   
   
	   



}

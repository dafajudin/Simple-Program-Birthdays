package main

import (
	"fmt"
)

func main(){
	fmt.Print("Input Your Birthday : ")
	var month int32
	fmt.Scan(&month)

	if month == 1 {
		fmt.Printf("Your Birthday is january")
	} else if month == 2 {
		fmt.Println("Your Birthday is February")
	} else if month == 3 {
		fmt.Println("Your Birthday is March")
	} else if month == 4 {
		fmt.Println("Your Birthday is April")
	} else if month == 5 {
		fmt.Println("Your Birthday is Mei")
	} else if month == 6{
		fmt.Println("Your Birthday is June")
	} else if month == 7{
		fmt.Println("Your Birthday is July")
	} else if month == 8{
		fmt.Println("Your Birthday is August")
	} else if month == 9{
		fmt.Println("Your Birthday is September")
	}else if month == 10{
		fmt.Println("Your Birthday is October")
	}else if month == 11{
		fmt.Println("Your Birthday is November")
	}else if month == 12{
		fmt.Println("Your Birthday is December")
	}	
}
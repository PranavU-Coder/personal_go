package main

import "fmt"

func main(){
	messageLen := 10
	maxmessageLen := 20

	if messageLen <= maxmessageLen {
		fmt.Println("message sent.")
	} else {
	 fmt.Println("not delivered")
	}
}

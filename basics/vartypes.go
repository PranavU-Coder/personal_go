package main

import "fmt"

func main() {
	const name = "PranavU"
	const openRate = 10.0

	// printf prints a formatted string to standard out
	// Sprintf returns a formatted string
	
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n",name,openRate)
	fmt.Print(msg)
}

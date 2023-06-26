package main

import "fmt"

func main() {
   
	/*
	old way -

	length := getLength(6)

	if length > 10 {
		fp("....")
	}
	just to save a line the below code is a syntax 
	*/

	if length := getLength(6) ; length > 10 {
		fmt.Println("the number is big")
	}

	fmt.Println("")
}


func getLength(number int) int {
	return number
}
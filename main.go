package main

import "fmt"

func main() {
    const name = "roch"
	const age = 22

	// const  msg = fmt.Sprintf("heya %v you're %v years old", name, age )
	/* Cannot do the above as const with sprintf throws an err */
	msg := fmt.Sprintf("heya %v you're %v years old", name, age )
	

	fmt.Println(msg)
}
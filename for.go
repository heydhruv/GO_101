package main

import "fmt"

// for is only construct for looping
func main() {
	// so no while loop but can be implemented like
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	for {
		fmt.Println("Hello")
	}

	//classic for loop
	// break and continue available

	for i:=0; i<5; i++ {
		fmt.Println(i)
	}

	// range
	for i:= range 3 {
		fmt.Println(i)
	}

}
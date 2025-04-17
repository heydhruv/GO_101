package main

import "fmt"

//by value
// func change(n int) {
// 	n = 5
// 	fmt.Println("in change", n)
// }

// by reference
func change(n *int) {
	//will change the global value
	*n = 5
	fmt.Println("in change", *n)
}

func main() {
	// n := 1
	// change(n)
	n := 1
	change(&n)
	fmt.Println(n)
}
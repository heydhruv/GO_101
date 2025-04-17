package main

import "fmt"

func counter() func() int {
	var count int = 0
	return func() int {
		count++
		return count
	}
}

func main() {
	inc:= counter()
	fmt.Println(inc()) // prints 1
	fmt.Println(inc())
}
package main

import "fmt"

func add(a int, b int) int {
	return a+b
}

//HOF
func process(check func(a int) int) {
	check(1)
}


func main() {
	s := add(5,5)
	fmt.Println(s)
	ch := func(a int)int{
		return a*2
	}
	process(ch)
}
package main

import "fmt"

func sum(nums...int)int {
	//can use interface{} for any type
	total:= 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	//variadic func can take n num of args (python *args)
	// in go its like spread operator of js
	s := sum(1,1,1)
	fmt.Println(s)

}
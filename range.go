package main

import "fmt"

func main() {
	// range can work like enumerate _ is index
	// for map us k,v
	nums := []int{1,2,3}
	sum := 0
	for _, num := range nums{
		sum = sum + num
	}
	fmt.Println(sum)

	// this smoker gives unicode lol
	// i is byte not index (rune)
	// typecast with string to get value
	for i, c := range "GO lang"{
		fmt.Println(i,c)
	}
}
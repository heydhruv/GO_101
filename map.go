package main

import (
	"fmt"
	"maps"
)

func main() {
	//object or dict
	// if key doesnt exist then returns default value of type (nil for any)
	m := make(map[string]any)
	m["name"] = "GO"
	fmt.Println(m["name"])
	fmt.Println(m["age"])
	// delete(m, "age")
	// clear(m)

	//map without make
	ma := map[string]any{"price":1}
	fmt.Println(ma["price"])

	//get value gives 2 elements value and ok which is bool
	value, ok := m["name"]
	fmt.Println(value, ok)

	fmt.Println(maps.Equal(m, ma))
}
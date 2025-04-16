package main

import "fmt"


func main() {
	age:= 18
	if age>= 18 {
		fmt.Println("Person is an adult")
	} else if age >=12 {
		fmt.Println("Person is a teenager")
	}else {
		fmt.Println("Person is kid")
	}
}
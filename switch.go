package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch
	i:= 3
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	case 3:
		fmt.Println("i is 3")
	default:
		fmt.Println("i is something else")
	}

	// multi conditional switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend baby")
	default:
		fmt.Println("it's weekday")
	}

	//type switch func
	whoImI := func(i interface{}) {
		switch t:= i.(type) {
			case int:
				fmt.Println("i is int")
			case string:
				fmt.Println("i is string")
			case bool:
				fmt.Println("i is bool")
			default:
				fmt.Println("i is something else", t)
		}
	}
	whoImI(1)
}
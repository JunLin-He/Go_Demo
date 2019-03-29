package main

import (
	"fmt"
)

func classchecker(items ...interface{}) { // Create a function which accept any numbers any type parameter
	for i, x := range items {
		switch x := x.(type) { // Create the shadow variable
		case bool:
			fmt.Printf("param #%d is a bool, value: %t\n", i, x)
		case float64:
			fmt.Printf("param #%d is a float64, value: %f\n", i, x)
		case int, int8, int16, int32, int64:
			fmt.Printf("param #%d is an int, value: %d\n", i, x)
		case uint, uint8, uint16, uint32, uint64:
			fmt.Printf("param #%d is a uint, value: %d\n", i, x)
		case nil:
			fmt.Printf("param #%d is a nil\n", i)
		case string:
			fmt.Printf("param #%d is a string, value: %s\n", i, x)
		default:
			fmt.Printf("param #%d's type is unknow\n", i)
		}
	}
}

func main() {
	classchecker(5, -17.98, "Alan", nil, true, complex(1, 1))
}

/*
	param #0 is an int, value: 5
	param #1 is a float64, value: -17.980000
	param #2 is a string, value: Alan
	param #3 is a nil
	param #4 is a bool, value: true
	param #5's type is unknow
*/

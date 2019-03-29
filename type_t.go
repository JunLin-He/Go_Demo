package main

import (
	"fmt"
)

func main() {
	x := uint16(6500)
	y := int16(x)                                           // Transform x to int16 type
	fmt.Printf("type and value of x is: %T and %d\n", x, x) // %T: Output the type of the variable x
	fmt.Printf("type and value of y is: %T and %d\n", y, y)

	var i interface{} = 99 // Create an interface{} type, value: 99
	var s interface{} = []string{"left", "right"}
	j := i.(int) // Assume i is compatiable int, and transform it to int by type assert
	fmt.Printf("type and value of j is: %T and %d\n", j, j)

	if s, ok := s.([]string); ok { // Create a shadow variable, the effect area cover the outside variable s
		fmt.Printf("%T -> %q\n", s, s)
	}
}

/*
	type and value of x is: uint16 and 6500
	type and value of y is: int16 and 6500
	type and value of j is: int and 99
	[]string -> ["left" "right"]
*/

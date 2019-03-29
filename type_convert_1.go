package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 65
	//b := string(a)
	b := strconv.Itoa(a)
	a, _ = strconv.Atoi(b)
	fmt.Println("b = ", b)
	fmt.Println("a = ", a)
}

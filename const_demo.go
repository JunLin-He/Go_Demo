package main

import (
	"fmt"
)

const (
	a = 'A'
	b = iota
	c = iota
	d
)

func main() {
	fmt.Println(a) // 65, 此时 iota的值为0
	fmt.Println(b) // 1
	fmt.Println(c) // 2
	fmt.Println(d) // 3
}

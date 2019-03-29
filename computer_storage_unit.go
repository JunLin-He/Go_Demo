package main

import (
	"fmt"
)

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println("B", B)
	fmt.Println("KB", KB)
	fmt.Println("MB", MB)
	fmt.Println("GB", GB)
	fmt.Println("TB", TB)
	//........
}

package main

import (
	"fmt"
)

func main() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("Func A")
}

func B() {
	/*
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Recover in B")
			}
		}()
	*/
	panic("Panic in B")
}

func C() {
	fmt.Println("Func C")
}

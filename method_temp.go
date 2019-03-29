package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{}
	a.Print()	
	fmt.Println(a.Name)
	
	b := B{}
	b.Print()
	fmt.Println(b.Name)
}

func (a *A) Print() {	// a is the reciever of the method
	a.Name = "AA"
	fmt.Println("A")
}

func (b B) Print()  {
	b.Name = "BB"
	fmt.Println("B")
}

//Output: 
// A
// AA
// B
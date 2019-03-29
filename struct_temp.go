package main 

import "fmt"

type human struct {
	Sex int
}

type teacher struct {
	human 
	Name string
	Age int 
	Sex int 
}

type student struct {
	human
	Name string
	Age int
}

func main() {
	a := teacher{Name: "Mary", Age: 30, human: human{Sex: 0}}
	b := student{Name: "Tom", Age: 13, human: human{Sex: 1}}
	a.Name = "Jenny"
	a.Age = 31
	a.human.Sex = 100
	a.Sex = 1
	fmt.Println(a, b)
}

// Output: {{100} Jenny 31} {{1} Tom 13}

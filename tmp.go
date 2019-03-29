package main

import (
	"fmt"
)

func main() {
	/*
		a := [...]int{99: 1}
		var p *[100]int = &a	// p为指向int100的数组的指针，且初始值为a的地址
		fmt.Println("%v", p)
	*/

	/*
		x, y := 1, 2
		a := [...]*int{&x, &y} // a为int100的指针的数组，即数组里装的数据都是整形指针
		fmt.Println(a)
	*/

	/*
		a := [2]int{1, 2}
		b := [2]int{1, 3}
		fmt.Println(a == b)		// true
	*/

	/*
		p := new([10]int)
		fmt.Println(p)
	*/

	/*
		a := [10]int{}
		a[1] = 2
		fmt.Println(a)
		p := new([10]int)
		p[1] = 2
		fmt.Println(p)
	*/

	a := [2][3]int{
		{1: 1},
		{2: 2}}
	fmt.Println(a) // Output: [[0 1 0] [0 0 2]]
}

package main

import "fmt"

func main() {
	/*
	   a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	   sa := a[2:5]
	   fmt.Println(string(sa))
	*/
	/*
		s1 := make([]int, 3)
		fmt.Println(len(s1), cap(s1))
	*/

	/*
		a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
		sa := a[2:5]
		fmt.Println(len(sa), cap(sa))
		sb := sa[1:3]
		fmt.Println(string(sb))
	*/

	/*
		s1 := make([]int, 3, 6)
		fmt.Printf("%p\n", s1)
		s1 = append(s1, 1, 2, 3)
		fmt.Printf("%v %p\n", s1, s1)
		s1 = append(s1, 1, 2, 3)
		fmt.Printf("%v %p\n", s1, s1)
	*/

	/*
		a := []int{1, 2, 3, 4, 5}
		s1 := a[2:5] // [3,4,5]
		s2 := a[1:3] // [2,3]
		fmt.Println(s1, s2)
		s2 = append(s2, 1, 2, 1, 1, 1, 1, 3, 2, 1, 3)
		s1[0] = 9
		fmt.Println(s1, s2)
	*/

	/*
		s1 := []int{1, 2, 3, 4, 5, 6}
		s2 := []int{7, 8, 9}
		copy(s1, s2) // copy s2 to s1
		fmt.Println(s1)
	*/
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9}
	copy(s2, s1) // copy s2 to s1
	fmt.Println(s2)
}

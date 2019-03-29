package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		m := make(map[int]string)
		m[1] = "OK"
		delete(m, 1) // delete function: delete the value of m[1], 1 is the key
		a := m[1]
		fmt.Println(a)
	*/

	/*
		var m map[int]map[int]string
		m = make(map[int]map[int]string)
		m[1] = make(map[int]string)
		m[2][1] = "OK"
		a := m[2][1]
		fmt.Println(a)
	*/

	/*
		var m map[int]map[int]map[int]string
		m = make(map[int]map[int]map[int]string)
		a, ok := m[1][1][1] // ok == true means the key-value exists
		if !ok {
			m[1] = make(map[int]map[int]string)
		}
		a, ok = m[1][1][1]
		if !ok {
			m[1][1] = make(map[int]string)
		}

		m[1][1][1] = "FOOD"
		a, ok = m[1][1][1]
		fmt.Println(a, ok)
	*/

	/*
		// i is the key or index (just like a counter), v is the value
		for i, v := range slice {

		}

		// the k and v is the backup, if you want to handle the map, you should use map[k]
		for k, v :=range map{
			map[k]
		}
	*/

	/*
		sm := make([]map[int]string, 5)
		fmt.Println(sm)
		for i := range sm {
			sm[i] = make(map[int]string, 1)
			sm[i][1] = "OK"
			fmt.Println(sm[i])
		}
		fmt.Println(sm)
	*/

	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	sort.Ints(s) // sort package: sort the key of s, s is integer
	fmt.Println(s)

}

package main

import (
	"fmt"
)

func main() {
	a := [...]int{5, 2, 6, 3, 9}
	fmt.Println(a)

	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}

/*
	a := [...]int{6, 3, 7, 10, 2}
	fmt.Println(a)

	num := len(a)
	var tmp int
	fmt.Println(num)

	for times := 0; times < (num - 2); times++ {
		for i := 0; i < (num - 1); i++ {
			if a[i] > a[i+1] {
				tmp = a[i+1]
				a[i+1] = a[i]
				a[i] = tmp
			}
		}
		fmt.Println(times)
	}

	fmt.Println(a)
*/

/*
   LABEL:
   	for i, _ := range a {
   		if a[i] > a[i+1] && i < (len(a)-1) {
   			a[i+1] = a[i]
   		}
   	}
   	num--
   	if num > 0 {
   		goto LABEL
   	}
*/

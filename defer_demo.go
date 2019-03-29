package main

import (
	"fmt"
)

func main() {
	var fs = [4]func(){} // equals to: var a = [10]int{}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() { fmt.Println("defer_closure i = ", i) }()
		fs[i] = func() { fmt.Println("closure i = ", i) }
	}

	for _, f := range fs {
		f()
	}
}

/* Output:
closure i = 0
closure i = 1
closure i = 2
closure i = 3
defer i = 4
defer_closure i = 4
defer i = 4
defer_closure i = 4
defer i = 4
defer_closure i = 4
defer i = 4
defer_closure i = 4
*/

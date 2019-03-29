package main

import (
	"fmt"
)

func main() {
	t0 := "\u4f55\u4fca\u9716" // t0: 何俊霖
	t1 := "\u559c\u6b22"       // t1: 喜欢
	t2 := "\u90d1\u70bc\u840d" // t2:  郑炼萍
	t3 := t0 + t1 + t2
	for index, char := range t3 {
		fmt.Printf("%-2d    %U    '%c'    %X    %d\n",
			index, char, char, []byte(string(char)), len([]byte(string(char))))
	}
	fmt.Printf("length of t0: %d, t1: %d, t2: %d\n", len(t0), len(t1), len(t2))
	fmt.Printf("content of t2[0:2] is: %X\n", t2[0:2])
}

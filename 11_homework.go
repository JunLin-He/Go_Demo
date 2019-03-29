package main 

import (
	"fmt"
)

type TZ int 

func (tz *TZ) Increase(num int)  {
	*tz += TZ(100)
}

func main() {
	var a TZ 
	a = 0
	a.Increase(100)
	fmt.Println(a)
}

/*
type TMP int

func main() {
	var a TMP
	a = 0
	a.Increase()
	fmt.Println(a)
}

func (a *TMP) Increase() {
	*a = *a + 100
}
*/
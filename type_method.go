package main 

import (
	"fmt"
)

type Tz int 

func main() {
	var a Tz
	a.Print()
	(*Tz).Print(&a)	
}

func (a *Tz) Print()  {
	fmt.Println("Tz")
}

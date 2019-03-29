package main

import (
	"fmt"
)

func main() {
	c := make(chan bool, 1)
	go func() {
		fmt.Println("Go Go Go!!!")
		//c <- true // 将值true存到channel c中，在c的值取出来之前，该channel一直处于阻塞的状态，只有运行完之后，才会通，然后main函数才会退出
		//close(c)
		<-c
	}()
	c <- true
	/*
		for v := range c {
			fmt.Println(v)
		}
	*/
}

package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)
	go func() {
		// 通过无限for循环与select组合来实现持续的接收与发送操作
		for {
			select {
			case v, ok := <-c1:
				if !ok { // 判断c1是否已经被关闭了
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)

			}
		}
	}()

	c1 <- 1
	c2 <- "Hi"
	c1 <- 5
	c2 <- "Hello"

	close(c1)

	for i := 0; i < 2; i++ {
		<-o
	}
	//<-o // c1与c2有任意一个关闭均会向o传一个值，而一旦读到o有值，main函数便会退出
}

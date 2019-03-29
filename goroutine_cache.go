package main

import (
	"fmt"
)

func main() {
	c := make(chan bool) // -> Go Go Go!!!
	go func() {
		fmt.Println("Go Go Go!!!")
		<-c
	}()
	c <- true
}

func main() {
	c := make(chan bool) // -> Go Go Go!!!
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true
	}()
	<-c
}

func main() {
	c := make(chan bool, 1) // -> Go Go Go!!!
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true
	}()
	<-c
}

func main() {
	c := make(chan bool, 1) // No output!
	go func() {
		fmt.Println("Go Go Go!!!")
		<-c
	}()
	c <- true
}

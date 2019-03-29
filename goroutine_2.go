package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}

	wg.Wait()
}

func Go(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	wg.Done()

}

/*
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}

	for i := 0; i < 10; i++ {
		<-c
	}
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	c <- true

}
*/
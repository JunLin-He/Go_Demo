package main

import (
	"fmt"
	"math/rand"
)

func main() {
	channels := make([]chan bool, 6) // Create a chan bool type slice, everyone is a channel that could sent bool value
	for i := range channels {        // Initialize the slice through 'range'
		channels[i] = make(chan bool)
	}

	go func() { // Run the anonymous function in other goroutine
		for {
			channels[rand.Intn(6)] <- true // rand.Intn(n int): create a random number that less than n
		} // send data to the random channel
	}()

	for i := 0; i < 36; i++ {
		var x int
		select { // select sentense: goto the branch which is not block
		case <-channels[0]:
			x = 1
		case <-channels[1]:
			x = 2
		case <-channels[2]:
			x = 3
		case <-channels[3]:
			x = 4
		case <-channels[4]:
			x = 5
		case <-channels[5]:
			x = 6
		}
		fmt.Printf("%d ", x)
	}
	fmt.Println()
}

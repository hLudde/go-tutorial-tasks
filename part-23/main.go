package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sayChannel(c chan int, integer int) {
	defer wg.Done()
	c <- integer
}

func main() {
	channel := make(chan int, 100)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go sayChannel(channel, i)
	}

	wg.Wait()
	close(channel)

	for item := range channel {
		fmt.Println(item)
	}
}

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Cleaned up: ", r)
	}
}

func say(text string, n int) {
	defer wg.Done()
	defer cleanup()
	for i := 0; i < n; i++ {
		fmt.Println(text)
		time.Sleep(100 * time.Millisecond)
		if i == 3 {
			panic("i cannot be 3!")
		}
	}
}

func finnishingStatement() {
	fmt.Println("And thats it folks!")
}

func main() {
	defer finnishingStatement()
	wg.Add(1)
	go say("Hello", 3)
	wg.Add(1)
	go say("World!", 5)
	wg.Wait()
	wg.Add(1)
	go say("These messages came after the previous go routines where done executing!", 50)
	wg.Wait()
}

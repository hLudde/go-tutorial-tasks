package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(text string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(text)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
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

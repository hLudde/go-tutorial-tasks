package main

import (
	"fmt"
	"time"
)

func say(text string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(text)
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	go say("Hello", 3)
	say("World!", 5)
}

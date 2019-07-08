package main

import (
	"fmt"
)

func sayChannel(c chan string, text string) {
	c <- text
}

func main() {
	channel := make(chan string)

	go sayChannel(channel, "world!")
	go sayChannel(channel, "Hello")

	text1, text2 := <-channel, <-channel

	fmt.Println(text1, text2)
}

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("The current time is: %d:%d:%d\n", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	fmt.Println("A random number between 0-100:", rand.Intn(100))
	fmt.Println("The squareroot of 100 is", sqrtOf100())
}

func sqrtOf100() float64 {
	return math.Sqrt(100)
}

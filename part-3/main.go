package main

import (
	"fmt"
)

func main() {
	variables(3.5, 2.7)
}

func variables(x, y float32) float32 {
	var i, j float32 = x, y
	fmt.Println(sum(i, j))
	fmt.Println(swap("Hello", "World!"))
	return 4.2
}

func sum(x, y float32) float32 {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}

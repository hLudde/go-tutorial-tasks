package main

import (
	"fmt"
)

func main() {
	x := 20
	y := &x
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(&x)
	fmt.Println(*y)

	*y = *y * *y

	fmt.Println(*&x)
}

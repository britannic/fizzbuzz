package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		var x interface{}
		switch {
		case i%3 == 0:
			x = "fizz"
		case i%5 == 0:
			x = "buzz"
		default:
			x = i
		}
		fmt.Println(x)
	}
}

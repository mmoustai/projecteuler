package main

import (
	"fmt"
)

// this is a comment
func fibonacci(c int) {
	sum, x, y := 0, 0, 1
	for i := 0; i < c; i += 1 {
		if x > 4000000 {
			fmt.Println("here")
			break
		}

		if x % 2 == 0{
		     sum += x
		}
		x, y = y, x+y
	}
	fmt.Println(sum)
}


func main() {
	fibonacci(500)
}

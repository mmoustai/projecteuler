package main

import (
	"fmt"
)

func isD(num int) bool {
	for j := 1; j <= 20; j += 1 {
		if num % j != 0 {
			return false
		}
	}
	return true
}


func main() {
	var num int
	i := true

	for i == true {
		num += 1
		if isD(num) {
			i = false
		}
	}

	fmt.Println(num)
}

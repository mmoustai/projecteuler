package main

import (
	"fmt"
	"strconv"
)

func isP(s string) bool {
	mid := len(s) / 2
	last := len(s) - 1
	for h := 0; h < mid; h++ {
		if s[h] != s[last-h] {
			return false
		}
	}
	return true
}

func main() {
	var l, x, y int
	for i := 1; i <= 999; i += 1 {
		for j := 1; j <= 999; j += 1 {
			var num = i * j
			if isP(strconv.Itoa(num)) {
				if num > l {
					l, x, y = num, i, j
				}

			}
		}

	}
	fmt.Println(l, x, y)
}

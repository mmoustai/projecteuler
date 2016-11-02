package main

import (
	"fmt"
)

func gcd(a, b uint64) uint64 {
	var remainder uint64 = 0
	for b != 0 {
		remainder = a % b
		a = b
		b = remainder
	}
	return a
}

func main() {
	var number, x_fixed, cycle_size, x, factor uint64 = 600851475143, 2, 2, 1000, 1
	var i uint64 = 0

	for factor == 1 {
		for i <= cycle_size && factor <= 1 {
			x = (x*x+1)%number
			factor = gcd(x - x_fixed, number)
		}

		cycle_size *= 2
		x_fixed = x
		i += 1
	}
	fmt.Println(factor)
}

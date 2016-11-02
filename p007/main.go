package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if (n%2 == 0) || (n%3 == 0) {
		return false
	}
	i := 5

	for i*i <= n {
		if (n%i == 0) || (n%(i+2) == 0) {
			return false
		}
		i += 6
	}

	return true
}

func main() {
	nbr, i := 0, 1
	var total int

	for nbr != 10001 {
		if isPrime(i) {
			nbr += 1
			total = i
		}
		i += 1
	}
	fmt.Println(total)
}

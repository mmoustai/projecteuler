package main

import (
	"fmt"
)

func main() {
	var sum, sum2 int

	for i := 0; i <= 100; i += 1 {
		sum += i
		sum2 += i * i
	}

	fmt.Println((sum * sum) - sum2)
}

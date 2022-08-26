package main

import (
	"fmt"
)

func main() {
	var N, count, a, min int
	fmt.Scanln(&N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&a)
		if i == 0 {
			min = a
		}
		if a < min {
			min = a
			count = 1
		} else if a == min {
			count++
		}
	}
	fmt.Print(count)
}

package main

import "fmt"

func main() {
	var a int
	var N int
	fmt.Scan(&N)

	array := make([]int, 0, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		if a > 0 {
			array = append(array, a)
		}
	}
	fmt.Print(len(array))
}

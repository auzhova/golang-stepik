package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	s := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&s[i])
	}
	fmt.Print(s[3])
}

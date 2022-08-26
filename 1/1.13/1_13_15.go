package main

import "fmt"

func main() {
	var A int
	fmt.Scan(&A)
	f1 := 0
	f2 := 1
	count := 0
	for f1 < A {
		f1, f2 = f2, f1 + f2
		count++
	}
	if f1 == A {
		fmt.Println(count)
	} else {
		fmt.Println(-1)
	}
}

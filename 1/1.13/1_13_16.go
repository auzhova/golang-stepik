package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	for i := 0; i < len(a); i++ {
		if a[i] == b[0] {
			continue

		} else {
			fmt.Print(string(a[i]))
		}
	}
}

package main

import "fmt"

func main() {
	var x int
	for fmt.Scan(&x); x <= 100; fmt.Scan(&x) {
		if x < 10 {
			continue
		}
		fmt.Println(x)
	}
}

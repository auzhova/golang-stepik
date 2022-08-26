package main

import "fmt"

func main() {
	var a int
	var max int = 0
	var count int = 0
	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		if max == a {
			count++
		} else if max < a {
			max = a
			count = 1
		}
	}
	fmt.Println(count)
}

package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	arr := []rune(s)
	var max rune = 0
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Print(string(max))
}

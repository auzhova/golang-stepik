package main

import "fmt"

func main() {
	var x, p, y int
	fmt.Scan(&x, &p, &y)
	var i int
	for i = 0; x <= y; i++ {
		x = x + (x * p / 100)
	}
	fmt.Println(i)
}

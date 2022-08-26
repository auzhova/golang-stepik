package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	a := x / 100
	b := x / 10 % 10
	c := x % 10
	fmt.Print(a + b + c)
}

package main

import "fmt"

func main() {
	var x int
	var a int
	var b int
	var c int
	var d int
	var e int
	var f int
	fmt.Scan(&x)
	a = x / 100000
	b = x / 10000 % 10
	c = x / 1000 % 10
	d = x / 100 % 10
	e = x / 10 % 10
	f = x % 10
	if a+b+c == d+e+f {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

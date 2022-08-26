package main

import "fmt"

func main() {
	var x int
	var a int
	var b int
	var c int
	fmt.Scan(&x)
	a = x % 400
	b = x % 4
	c = x % 100
	if a == 0 || (b == 0 && c != 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

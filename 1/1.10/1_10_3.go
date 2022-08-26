package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	var sum int = 0
	for i := a; i <= b; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}

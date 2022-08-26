package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)
	h := k / 3600
	m := k % 3600 / 60
	fmt.Println("It is", h, "hours", m, "minutes.")
}

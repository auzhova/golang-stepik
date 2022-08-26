package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	arr := []rune(s)
	for _, value := range arr {
		x, _ := strconv.Atoi(string(value))
		fmt.Print(x * x)
	}
}
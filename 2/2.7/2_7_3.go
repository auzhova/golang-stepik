package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string
	var str []string
	fmt.Scan(&text)
	array := []rune(text)

	for _, value := range array {
		str = append(str, string(value))
	}
	fmt.Println(strings.Join(str, "*"))
}

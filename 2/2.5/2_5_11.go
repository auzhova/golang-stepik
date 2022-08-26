package main

import (
	"fmt"
	"strings"
)

func main() {
	var str, result string
	fmt.Scan(&str)
	result = ""
	for i := 0; i < len(str); i++ {
		if strings.Count(str, string(str[i])) > 1 {
			continue
		}
		result += string(str[i])
	}
	fmt.Print(result)
}

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var str string
	fmt.Scan(&str)
	rune := []rune(str)
	if len(rune) < 5 {
		fmt.Print("Wrong password")
		return
	}
	for i := range rune {
		if !unicode.Is(unicode.Latin, rune[i]) && !unicode.IsDigit(rune[i]) {
			fmt.Print("Wrong password")
			return
		}
	}
	fmt.Print("Ok")
}

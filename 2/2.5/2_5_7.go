package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	value := []rune(text)

	if unicode.IsUpper(value[0]) && value[len(value)-1] == '.' {
		fmt.Print("Right")
	} else {
		fmt.Print("Wrong")
	}
}

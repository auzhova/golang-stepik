package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text := []rune(input)
	str := ""
	for i := 1; i < len(text); i++ {
		if i % 2 == 1 {
			str = str + string(text[i])
		}
	}
	fmt.Print(str)
}
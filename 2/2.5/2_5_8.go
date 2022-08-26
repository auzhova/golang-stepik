package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	value := []rune(text)
	var newText string
	l := len(value) - 1

	for i := l; i >= 0; i-- {
		newText += string(value[i])
	}

	if newText == text {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}

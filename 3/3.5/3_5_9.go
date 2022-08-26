package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	// Все указанные в тексте задачи пакеты уже импортированы (strconv, bufio, os, io)
	// Другие использовать нельзя.
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		x, _ := strconv.Atoi(str)
		sum += x
	}

	io.WriteString(os.Stdout, strconv.Itoa(sum))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := scanner.Text()
	date, err := time.Parse("2006-01-02 15:04:05", data)

	if err != nil {
		panic(err)
	}

	if date.Hour() >= 13 {
		date = date.Add(24 * time.Hour)
	}
	fmt.Println(date.Format("2006-01-02 15:04:05"))
}

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case a < 10:
		fmt.Print(a)
	case a < 100:
		fmt.Print(a / 10)
	case a < 1000:
		fmt.Print(a / 100)
	case a < 10000:
		fmt.Print(a / 1000)
	case a < 100000:
		fmt.Print(a / 10000)
	}
}

package main

import "fmt"

func main() {
	var n int
	var word string
	fmt.Scan(&n)

	if n == 11 {
		word = "korov"
	} else if n % 10 == 1 {
		word = "korova"
	} else if (n < 10 || n > 20) && (n % 10 == 2 || n % 10 == 3 || n % 10 == 4) {
		word = "korovy"
	} else {
		word = "korov"
	}

	fmt.Printf("%d %s", n, word)
}
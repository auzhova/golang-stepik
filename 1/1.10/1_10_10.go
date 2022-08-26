package main

import "fmt"

func main() {
	var s string = ""
	var a, b int
	fmt.Scan(&a, &b)
	for a > 0 {
		c := b
		d := a % 10
		a = a / 10
		for c > 0 {
			if c%10 == d {
				s = fmt.Sprint(d) + " " + s
			}
			c = c / 10
		}
	}
	fmt.Println(s)
}

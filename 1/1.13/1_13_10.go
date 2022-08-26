package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n)
	sum = 0
	for {
		sum += n % 10
		n = n / 10
		if n < 10 {
			sum += n
			if sum < 10 {
				break
			} else {
				n = sum
				sum = 0
			}
		}
	}
	fmt.Print(sum)
}
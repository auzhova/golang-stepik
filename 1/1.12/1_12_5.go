package main

import "fmt"

func main() {
	var workArray [10]uint8
	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}

	var x, y int
	for n := 0; n < 3; n++ {
		fmt.Scan(&x, &y)
		workArray[x], workArray[y] = workArray[y], workArray[x]
	}
	for i := 0; i < 10; i++ {
		fmt.Print(workArray[i], " ")
	}
}

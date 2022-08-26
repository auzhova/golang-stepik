package main

import "fmt"

func main() {
	var a int
	var h int
	var m int
	fmt.Scan(&a)
	// Если 360 градусов - это 12 часов (или 12 * 60 = 720 минут), то 1 градус - это две минуты,
	// 30 градусов - это один час, поэтому количество целых часов будет = d div 30,
	// а минуты получаются из формулы, по которой подсчитывается остаток: d mod 30.
	// И этот результат надо умножить на 2, так как один градус - это 2 минтуы.
	h = a / 30
	m = a % 30 * 2
	fmt.Println("It is", h, "hours", m, "minutes.")
}
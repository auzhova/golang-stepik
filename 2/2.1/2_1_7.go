func minimumFromFour() int {
	var a, b, c, d, min int
	fmt.Scan(&a, &b, &c, &d)
	arr := [4]int{a, b, c, d}

	for i := 0; i < 4; i++ {
	fmt.Scan(&arr[i])
	if i == 0 {
		min = arr[i]
	}

	if min > arr[i] {
		min = arr[i]
	}
	}
	return min
}
func sumInt(array ...int) (int, int) {
	count := len(array)
	sum := 0

	for i := 0; i < count; i++ {
		sum = sum + array[i]
	}

	return count, sum
}
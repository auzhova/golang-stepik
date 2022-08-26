func fibonacci(n int) int {
	f1 := 0
	f2 := 1

	for i := 1; i < n; i++ {
		f1, f2 = f2, f1+f2
	}

	return f2
}
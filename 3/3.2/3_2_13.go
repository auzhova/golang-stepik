func adding(a string, b string) int64 {
	r1 := []rune(a)
	r2 := []rune(b)
	strA := ""
	strB := ""
	for i := range r1 {
		if unicode.IsDigit(r1[i]) {
			strA += string(r1[i])
		}
	}
	for i := range r2 {
		if unicode.IsDigit(r2[i]) {
			strB += string(r2[i])
		}
	}

	x, _ := strconv.Atoi(strA)
	y, _ := strconv.Atoi(strB)

	return int64(x + y)
}
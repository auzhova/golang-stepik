fn := func(x uint) uint {
	var s string
	str := strconv.Itoa(int(x))

	for _, value := range []rune(str) {
		num, _ := strconv.Atoi(string(value))
		if num % 2 == 0 && num != 0 {
			s += strconv.Itoa(num)
		}
	}

	res, _ := strconv.Atoi(s)

	if res == 0 {
		return 100
	}

	return uint(res)
}
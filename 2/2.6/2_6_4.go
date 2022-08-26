func main() {
	var a, b int
	fmt.Scan(&a, &b)
	res, err := divide(a, b)

	if err != nil {
		fmt.Print("ошибка")
	} else {
		fmt.Print(res)
	}
}
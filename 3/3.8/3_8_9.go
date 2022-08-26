func task2(channel chan string, str string) {
	for i := 0; i < 5; i++ {
		channel <- str + " "
	}
}
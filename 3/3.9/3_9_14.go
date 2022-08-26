func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	channel := make(chan int)
	sum := 0

	go func() {
		var sum int
		defer close(channel)
		for {
			select {
				case arg := <-arguments:
					sum += arg
				case <-done:
					channel <- sum
				return
			}
		}
	}()
	return channel
}
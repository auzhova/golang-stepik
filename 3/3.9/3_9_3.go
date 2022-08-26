done := make(chan struct{})
	go func() {
		work()
		close(done)
	}()
<-done
}

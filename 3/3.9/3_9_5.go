wg := new(sync.WaitGroup)

for i := 0; i < 10; i++ {
	wg.Add(1) // Увеличиваем счетчик горутин в группе
	go func() {
		defer wg.Done()
		work()
	}()
}
wg.Wait() // ожидаем завершения всех горутин в группе
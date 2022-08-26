// Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!
func task(channel chan int, N int) {
	channel <- N + 1
}
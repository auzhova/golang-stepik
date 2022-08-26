func removeDuplicates(inputStream chan string, outputStream chan string){
	var prev string
	for cur := range inputStream {
		if prev != cur {
			outputStream <- cur
		}
			prev = cur
		}
	close(outputStream)
}
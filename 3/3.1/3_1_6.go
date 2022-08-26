for _, city := range groupCity[10] {
	if cityPopulation[city] != 0 {
		delete(cityPopulation, city)
	}
}

for _, city := range groupCity[1000] {
	if cityPopulation[city] != 0 {
		delete(cityPopulation, city)
	}
}
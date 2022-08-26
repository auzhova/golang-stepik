package main

import (
	"fmt"
	"time"
)

// вам это понадобится
const now = 1589570165

func main() {
	var min, sec string

	fmt.Scanf("%s мин. %s сек.", &min, &sec)

	duration, err := time.ParseDuration(fmt.Sprintf("%sm%ss", min, sec))
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Unix(now, 0).Add(duration).UTC().Format(time.UnixDate))
}

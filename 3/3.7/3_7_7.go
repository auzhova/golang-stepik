package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	arr := strings.Split(in.Text(), ",")
	Time1, err1 := time.Parse("02.01.2006 15:04:05", arr[0])
	if err1 != nil {
		panic(err1)
	}
	Time2, err2 := time.Parse("02.01.2006 15:04:05", arr[1])
	if err2 != nil {
		panic(err2)
	}
	if Time1.Unix() >= Time2.Unix() {
		fmt.Println(Time1.Sub(Time2))
	} else {
		fmt.Println(Time2.Sub(Time1))
	}
}

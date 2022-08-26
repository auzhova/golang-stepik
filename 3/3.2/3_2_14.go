package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Replace(str, string('\n'), "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, ",", ".", -1)
	arr := strings.Split(str, ";")
	x, _ := strconv.ParseFloat(arr[0], 64)
	y, _ := strconv.ParseFloat(arr[1], 64)

	fmt.Printf("%.4f", x/y)
}

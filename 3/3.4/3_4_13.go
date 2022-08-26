package main

import (
	"fmt" // пакет используется для проверки ответа, не удаляйте его
	"strings"
)

type Battery struct {
	Charge string
}

func (b Battery) String() string {
	return fmt.Sprintf("[%v]", b.Charge)
}

func main() {
	var s string
	fmt.Scanln(&s)
	count0 := strings.Count(s, "0")
	count1 := strings.Count(s, "1")
	res0 := strings.Repeat(" ", count0)
	res1 := strings.Repeat("X", count1)

	batteryForTest := Battery{res0 + res1}
}

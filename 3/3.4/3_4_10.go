package main

import (
	"encoding/json" // пакет используется для проверки ответа, не удаляйте его
	"fmt"           // пакет используется для проверки ответа, не удаляйте его
	"os"            // пакет используется для проверки ответа, не удаляйте его
)

func checkValue(value interface{}) bool {
	switch v := value.(type) {
	case float64:
		return true
	default:
		fmt.Printf("value=%v: %T", v, v)
		return false
	}
}

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
											// все полученные значения имеют тип пустого интерфейса

	if checkValue(value1) && checkValue(value2) {
		switch operation.(type) {
		case string:
			switch operation {
			case "+":
				fmt.Printf("%.4f", value1.(float64)+value2.(float64))
			case "-":
				fmt.Printf("%.4f", value1.(float64)-value2.(float64))
			case "/":
				fmt.Printf("%.4f", value1.(float64)/value2.(float64))
			case "*":
				fmt.Printf("%.4f", value1.(float64)*value2.(float64))
			}
		default:
			fmt.Println("неизвестная операция")
			return
		}
	}
}
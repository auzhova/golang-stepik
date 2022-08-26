package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Average struct {
	Average float32
}

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type Data struct {
	Id       int
	Number   string
	Year     int
	Students []Student
}

func main() {
	var count int = 0
	var data Data

	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(stdin, &data); err != nil {
		panic(err)
	}

	for _, student := range data.Students {
		count += len(student.Rating)
	}

	var average Average = Average{float32(count) / float32(len(data.Students))}

	result, err := json.MarshalIndent(average, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))
}

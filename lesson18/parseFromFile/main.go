package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Student struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Age      int    `json:"age"`
}

type StudentsData struct {
	Students []Student `json:"students"`
}

func main() {
	var students []StudentsData
	///////////////////////////////////////////////////////
	file, err := os.Open("temp.json")
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	/////////////////////////////////////////////////////
	// fmt.Println(string(data))

	err = json.Unmarshal(data, &students)
	fmt.Println(students)

}

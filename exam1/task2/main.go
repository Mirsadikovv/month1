package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

/*
2. employees.json filedan hamma employees larni o'qib olib. orderByAge.txt filega age
bo'yicha sort(kichkinadan kattaga) qilib yozilsin. Top 3ta oyligi eng yuqori employeeni
topSalaryEmployees.txt filega yozilsin.
2ta filega ham yozilish formati(
"first_name": "Adam" "second_name": "Johns" "salary": 2500 "experience": 3 "age": 21
"first_name": "Alice" "second_name": "Smith" "salary": 2800 "experience": 5 "age": 25
)
*/
type Employees struct {
	First_name  string `json:"first_name"`
	Second_name string `json:"second_name"`
	Salary      int    `json:"salary"`
	Experience  int    `json:"experience"`
	Age         int    `json:"age"`
}

func main() {
	var str string
	emp := []Employees{}
	file, err := os.Open("employees.json")
	if err != nil {
		panic(err)
	}

	dataJson, err := io.ReadAll(file)
	// fmt.Println(string(dataJson))

	err = json.Unmarshal(dataJson, &emp)
	if err != nil {
		panic(err)
	}

	// fmt.Println(emp)

	orderByAge, err := os.Create("orgerByAge.txt")
	if err != nil {
		panic(err)
	}
	topSalaryEmployees, err := os.Create("topSalaryEmployees.txt")
	if err != nil {
		panic(err)
	}

	sort.Slice(emp, func(i, j int) bool {
		return emp[i].Age < emp[j].Age
	})
	for _, val := range emp {

		str = fmt.Sprint(" \"first_name\" : ", val.First_name, ", \"second_name\" : ", val.Second_name,
			", \"salary\" : ", val.Salary, ", \"experience\" : ", val.Experience, ", \"age\" : ", val.Age, "\n")
		_, err = io.WriteString(orderByAge, str)
		if err != nil {
			panic(err)
		}

	}
	sort.Slice(emp, func(i, j int) bool {
		return emp[i].Salary > emp[j].Salary
	})

	for i, val := range emp {
		str = fmt.Sprint(" \"first_name\" : ", val.First_name, ", \"second_name\" : ", val.Second_name,
			", \"salary\" : ", val.Salary, ", \"experience\" : ", val.Experience, ", \"age\" : ", val.Age, "\n")
		_, err = io.WriteString(topSalaryEmployees, str)
		if err != nil {
			panic(err)
		}
		if i == 2 {
			break
		}

	}
}

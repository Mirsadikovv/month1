package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name     string          `json:"name"`
	LastName string          `json:"last_name"`
	Age      int             `json:"age"`
	Subjects json.RawMessage `json:"subjects" `
	Grades   []struct {
		Subject string `json:"subject"`
		Grade   int    `json:"grade"`
	} `json:"grades"`
}

func main() {
	var student interface{}

	newJson := `{
		"name": "Student",
		"last_name": "Familya",
		"age": 2,
		"grade": {},
		"subject":["Student"]
	}`

	err := json.Unmarshal([]byte(newJson), &student)
	if err != nil {
		panic(err)
	}

	newStudent := student.(map[string]interface{})

	for key, value := range newStudent {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}

	fmt.Printf("%T \n", newStudent)

	fmt.Println("student: ", student.(map[string]interface{}))

	// newJson := `{"name":"informatika,code:123, "grade":5, }`

	// jsonData1, err := json.Marshal(newJson)
	// if err != nil {
	// 	panic(err)
	// }

	// var student = Student{
	// 	LastName: "Familya",
	// 	Name:     "Dilshod",
	// 	Age:      1000,
	// 	Subjects: json.RawMessage(jsonData1),
	// 	// Grades: []struct{
	// 	// 	{Subject: "Algebra", Grade: 3},
	// 	// },
	// }

	// jsonData, err := json.Marshal(student)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(jsonData))
}

package main

import (
	"fmt"
	"io"
	"os"
)

/*
5. Student degan struct name, grades(arrays of int),course. calculateAverage degan method
yozilsin u studentning o'rtacha grade ini hisoblaydi. Agar averageGrade 60dan kichkina
bo'lsa u failed bo'lgan hisoblanadi aksi bo'lsa passed. U fail yoki passed bo'lgani haqida
result.txt filega yo'zing. (You are failed yoki You are passed)
*/

type Student struct {
	Name   string
	Grades []int
	course int
}

func (s *Student) calculateAverage() float64 {
	var sum, avg float64
	for _, val := range s.Grades {
		sum += float64(val)
	}
	avg = ((sum / float64(len(s.Grades))) / 5) * 100
	return avg
}

func main() {
	student := Student{Name: "Noname", Grades: []int{2, 3, 4, 2, 5, 5, 5}, course: 3}
	file, err := os.Create("result.txt")
	if err != nil {
		panic(err)
	}
	score := student.calculateAverage()
	if score < 60 {
		_, err = io.WriteString(file, fmt.Sprint("You are failed :(\nYour score -> ", score))
		if err != nil {
			panic(err)
		}
	} else {
		_, err = io.WriteString(file, fmt.Sprint("You are passed congrats!!!\nYour score -> ", score))
		if err != nil {
			panic(err)
		}
	}

}

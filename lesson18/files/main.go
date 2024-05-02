package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file, err := os.Open("temp.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// ans, err := os.ReadFile("temp.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(ans))
	// err = os.WriteFile("temp.txt", []byte("sd123"), 0644)
	// if err != nil {
	// 	panic(err)
	// }
	// file, err := os.OpenFile("temp.txt", os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// _, err = file.WriteString("b")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // fmt.Println(n, 12312312312)

	// ans1, err := os.ReadFile("temp.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(ans1))

	// err = os.Rename(".env", "temp.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// info, err := os.Stat("temp.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(info.IsDir())
	// fmt.Println(os.Getwd())
	// err = os.Setenv("ip", "localhos")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(os.Getenv("ip"))

	/////////////////////////////////////////////////////////////////

	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла для чтения:", err)
		return
	}
	defer inputFile.Close()

	// Открываем файл для записи
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла для записи:", err)
		return
	}
	defer outputFile.Close()

	// Читаем все данные из файла
	data, err := io.ReadAll(inputFile)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	// Записываем считанные данные в другой файл
	_, err = io.WriteString(outputFile, string(data))
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	fmt.Println("Данные успешно скопированы из input.txt в output.txt")
}

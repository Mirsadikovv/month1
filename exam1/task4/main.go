package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

/*
4. file yaratish kerak.Terminaldan fileName va recomendationCount kiritiladi. Agar
fileName daka file bo'masa yaratib qo'yadi.Agar unaqa file bo'sa manga
recomendationCountga nechi kiritilgan bo'lsa o'shancha name rekommend(recommend logikasi
fileName+randomraqam) qilishi kerak. Shunaqa namelik file yarata olasan deb.Pragramma 1ta
file yaratmaguncha ishlashi kerak.
*/

func main() {
	var fileName, nstr, countStr string
	var count, n, i int
	names := allFiles()

	for {
		fmt.Println("filename:")
		fmt.Scan(&fileName)
		fmt.Println("count:")
		fmt.Scan(&count)
		countStr = strconv.Itoa(count)
		str := fmt.Sprint(fileName + countStr)
		if check(names, str) {
			fmt.Println("попробуйте следующие варианты:")
			for i != count {
				n = rand.Intn(100)
				nstr = strconv.Itoa(n)
				fmt.Println(fmt.Sprint(fileName + nstr))
				i++

			}
		} else {
			_, err := os.Create(str)
			if err != nil {
				panic(err)
			}
			break
		}

	}
}

func allFiles() (names []string) {

	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		names = append(names, file.Name())
	}
	return names
}

func check(slice []string, str string) bool {
	for _, val := range slice {
		if val == str {
			return true
		}
	}
	return false
}

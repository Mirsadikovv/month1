package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
1. data.txt filedan hamma contentlarni o'qib olib sonlarni(string ko'rinishida)
numbers.txt filega new line bilan yozish va eng ohirgi lineda sonlarni jami
summasini yozish. So'zlarni qaysiki faqat A yoki O yoki a yoki o harflari qatnashgan
so'zlarni alohida words.txt filega yozish new line bilan
*/

func main() {
	var sum, n int
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	words, err := os.Create("words.txt")
	if err != nil {
		panic(err)
	}
	numbers, err := os.Create("numbers.txt")

	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(file)
	data1 := string(data)
	data2 := strings.Split(string(data1), " ")
	for _, word := range data2 {
		if strings.Contains(string(word), "o") || strings.Contains(string(word), "a") || strings.Contains(string(word), "A") || strings.Contains(string(word), "O") {
			_, err = io.WriteString(words, string(word)+"\n")
			if err != nil {
				panic(err)
			}
		}
		for _, let := range word {
			if unicode.IsDigit(let) {
				_, err = io.WriteString(numbers, string(let))
				if err != nil {
					panic(err)
				}
			}

		}

		if strings.Contains(word, "0") || strings.Contains(word, "1") ||
			strings.Contains(word, "2") || strings.Contains(word, "3") ||
			strings.Contains(word, "4") || strings.Contains(word, "5") ||
			strings.Contains(word, "6") || strings.Contains(word, "7") ||
			strings.Contains(word, "8") || strings.Contains(word, "9") {
			_, err = io.WriteString(numbers, "\n")
			if err != nil {
				panic(err)
			}
		}
	}

	nums, err := os.Open("numbers.txt")
	defer nums.Close()
	numsData, err := io.ReadAll(nums)
	if err != nil {
		panic(err)
	}
	numsStr := strings.Split(string(numsData), "\n")
	for _, v := range numsStr[len(numsStr)-2] {
		n, err = strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}
		sum += n
	}
	fmt.Println("sum of last line - >", sum)

}

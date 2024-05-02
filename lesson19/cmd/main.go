package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// intToBin()
	// floatToInt()
	// hexToInt()

	// n := time.Now()
	// nn := time.Since(n)
	// fmt.Println(generateRandomNumber(), nn)

	///////////////////////////////////////////////////
	// z := 12
	// add := func(x, y int) int {
	// 	return (x + y) * z
	// }
	// fmt.Println(add(2, 3), " closure")
	///////////////////////////////////////////////

	/////////////////////////////////////////////////////
	// c := 0
	// count := func() int {
	// 	c++
	// 	return c
	// }
	// count()
	// count()
	// count()

	// fmt.Println(count(), " closure")
	////////////////////////////////////////////////////
	arrayToSlice()
}

func intToBin() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter integer number")
	intInput, _ := reader.ReadString('\n')
	intInput = intInput[:len(intInput)-1]
	intValue, err := strconv.Atoi(intInput)
	if err != nil {
		panic(err)
	}
	binaryData := strconv.FormatInt(int64(intValue), 2)
	fmt.Println("binaryData - ", binaryData)
}

func floatToInt() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter float number - ")
	floatInput, _ := reader.ReadString('\n')
	floatInput = floatInput[:len(floatInput)-1]
	floatVal, err := strconv.ParseFloat(floatInput, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("intData - ", int(floatVal))
}

func hexToInt() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter hexadecimal number - ")
	hexInput, _ := reader.ReadString('\n')
	hexInput = hexInput[2 : len(hexInput)-1]
	hexVal, err := strconv.ParseInt(hexInput, 16, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("intData - ", int(hexVal))
}

func isCapitalThere(str string) bool {
	for _, val := range str {
		if string(val) >= "A" && string(val) <= "Z" {
			fmt.Println(val)
			return true
		}
	}
	return false
}

func isDigitThere(str string) bool {
	for _, val := range str {
		if string(val) > "0" && string(val) < "9" {
			return true
		}
	}
	return false
}

func isSymbolThere(str string) bool {
	for _, val := range str {
		if strings.Contains("!@#$%^&*", string(val)) {
			return true
		}
	}
	return false
}

func isPalindrome(str string) bool {
	start := 0
	end := len(str) - 1

	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func countAlpha(str string) int {
	var c int
	for _, val := range str {
		if string(val) >= "a" && string(val) <= "z" || string(val) >= "A" && string(val) <= "Z" {
			c++
		}
	}
	return c
}

func generateRandomNumber() int {
	a := rand.Intn(1000)
	return a
}

func CalculateGrade(s int) string {
	str := ""
	switch {
	case s > 89:
		str = "A"

	case s > 79:
		str = "B"

	case s > 69:
		str = "C"
	case s > 59:
		str = "D"
	default:
		str = "F"
	}

	return str
}

func arrayToSlice() {
	var arr [5]int
	var val, sum int
	for i := 0; i < 5; i++ {
		val = rand.Intn(100)
		arr[i] = val
	}
	fmt.Println("array -> ", arr)
	slice := arr[:]
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	for _, val := range slice {
		sum += val
	}
	fmt.Println("sortedSlice -> ", slice, "\n", "sum ->", sum, "\n", "min ->", slice[0], "\n", "max ->", slice[len(slice)-1])
}

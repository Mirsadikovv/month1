package main

import "fmt"

func incrementValues(list []int, a int) []int {
	for i, val := range list {
		list[i] = val + a
	}
	return list
}

func printValues(slice []int) {
	fmt.Println(slice)
}

func incrementValuesPtr(list *[]int, a int) []int {
	for i, val := range *list {
		(*list)[i] = val + a
	}
	return *list
}
func main() {
	slice := []int{1, 2, 3, 4}
	printValues(slice)
	incrementValues(slice, 5)
	printValues(slice)
	incrementValuesPtr(&slice, 12)
	printValues(slice)
}

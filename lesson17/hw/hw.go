package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Koment string `json:"koment"`
}

func main() {
	var ord interface{}
	order := Order{
		Id:     12,
		Name:   "zoloto",
		Price:  560,
		Koment: "dorogo",
	}
	jsonData, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))

	json.Unmarshal([]byte(jsonData), &ord)
	fmt.Println(ord)
}

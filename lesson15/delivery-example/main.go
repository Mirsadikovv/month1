package main

import (
	"fmt"
	"lesson15/courier"
	"lesson15/customer"
	"lesson15/order"
	"lesson15/storage"
	"strconv"
	"time"
)

func main() {
	// var customer customer.Customer

	newCustomer := customer.NewCustomer(4, "name", "address", "phone")

	// fmt.Println("newCustomer", newCustomer)

	newOrder := order.Order{
		CustomerID:  newCustomer.ID,
		ProductName: "product",
		Amount:      9,
	}

	order := order.NewOrder(newOrder)
	fmt.Println("newOrder", order)

	courier := courier.Courier{}

	courier.NewCourier("name", "Nexia 2", "avto")

	// fmt.Println("courier", courier)

	storage := storage.NewCache()

	storage.Set(strconv.Itoa(newCustomer.ID), newCustomer, 5*time.Minute)

	data, err := storage.Get(strconv.Itoa(newCustomer.ID))
	if err != nil {
		panic(err)
	}
	fmt.Println("data", data)

	newCustomer1 := customer.NewCustomer(5, "name12332122", "address1", "phone1")
	fmt.Println(data)
	storage.Update(strconv.Itoa(newCustomer.ID), newCustomer1)

	data1, err1 := storage.Get(strconv.Itoa(newCustomer.ID))
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("data", data1)
	// storage.Delete("kuryer")

	// data, err := storage.Get("kuryer")
	// if err != nil {
	// 	panic(err)
	// }

}

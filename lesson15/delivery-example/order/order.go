package order

type Order struct {
	CustomerID  int
	ProductName string
	Amount      int
}

func NewOrder(givenOrder Order) Order {
	return givenOrder
}

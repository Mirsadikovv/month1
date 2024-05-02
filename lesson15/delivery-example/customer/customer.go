package customer

type Customer struct {
	ID      int
	Name    string
	Address string
	Phone   string
}

func NewCustomer(id int, name, address, phone string) *Customer {
	return &Customer{
		ID:      id,
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}

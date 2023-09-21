package postgre

import "fmt"

type OrderRepository struct {
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (r *OrderRepository) Create(order string) {
	fmt.Printf("Order %s created in postgresql \n", order)
}

func (r *OrderRepository) Update(order string, newName string) {
	fmt.Printf("Order %s updated in postgresql \n", order)
}

func (r *OrderRepository) Delete(order string) {
	fmt.Printf("Order %s deleted in postgresql \n", order)
}

func (r *OrderRepository) IsMatch(dbType string) bool {
	return dbType == "postgre"
}

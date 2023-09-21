package mongo

import "fmt"

type OrderRepository struct {
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (r *OrderRepository) Create(order string) {
	fmt.Printf("Order %s created in mongodb\n", order)
}

func (r *OrderRepository) Update(order string, newName string) {
	fmt.Printf("Order %s updated in mongodb\n", order)
}

func (r *OrderRepository) Delete(order string) {
	fmt.Printf("Order %s deleted in mongodb\n", order)
}

func (r *OrderRepository) IsMatch(dbType string) bool {
	return dbType == "mongo"
}

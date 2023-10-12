package repository

import (
	"github.com/alibekabdrakhman/justcode/lecture2/third_task/repository/mongo"
	"github.com/alibekabdrakhman/justcode/lecture2/third_task/repository/postgre"
)

type IOrderRepository interface {
	Create(order string)
	Update(order string, newName string)
	Delete(order string)
	IsMatch(dbType string) bool
}

type Storage struct {
	Order IOrderRepository
}

func NewStorage(dbType string) *Storage {
	if dbType == "mongo" {
		return &Storage{
			Order: mongo.NewOrderRepository(),
		}
	}
	return &Storage{
		Order: postgre.NewOrderRepository(),
	}
}

package main

import "github.com/alibekabdrakhman/justcode/lecture2/third_task/repository"

func main() {
	mongo := repository.NewStorage("mongo")
	mongo.Order.Create("Order")
	mongo.Order.Update("Order", "Order1")
	mongo.Order.Delete("Order")

	postgre := repository.NewStorage("postgre")
	postgre.Order.Create("Order2")
	postgre.Order.Update("Order2", "Order3")
	postgre.Order.Delete("Order2")
}

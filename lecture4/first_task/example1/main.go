package main

import (
	"fmt"
	"sync"
)

type CreditCard struct {
	balance int
	mutex   sync.Mutex
}

func (c *CreditCard) put(amount int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.balance += amount
}
func (c *CreditCard) get(amount int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.balance >= amount {
		c.balance -= amount
	}
}
func main() {
	card := CreditCard{balance: 1000}
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 8; i++ {
				fmt.Println(card.balance)

				card.get(5)
			}
		}()
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 8; i++ {
				fmt.Println(card.balance)
				card.put(5)
			}
		}()
	}
	wg.Wait()

	fmt.Println(card.balance)
}

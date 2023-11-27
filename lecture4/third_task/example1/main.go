package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var list []int
	var mutex sync.RWMutex
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			value := rand.Intn(100)
			mutex.Lock()
			defer mutex.Unlock()
			list = append(list, value)
			fmt.Println("Wrote: ", value)
		}(i)
	}
	time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mutex.RLock()
			defer mutex.RUnlock()
			if len(list) > 0 {
				index := rand.Intn(len(list))
				value := list[index]
				fmt.Println("Read: ", value)
			} else {
				fmt.Println("List is empty")
			}
		}(i)
	}
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go func() {
		for _, num := range []int{1, 2, 3} {
			ch1 <- num
		}
		close(ch1)
	}()

	go func() {
		for _, num := range []int{4, 5, 6} {
			ch2 <- num
		}
		close(ch2)
	}()

	go func() {
		for _, num := range []int{7, 8, 9} {
			ch3 <- num
		}
		close(ch3)
	}()
	for num := range merge(ch1, ch2, ch3) {
		fmt.Println(num)
	}

}
func merge(channels ...<-chan int) <-chan int {
	merged := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		for _, ch := range channels {
			wg.Add(1)
			go func(ch <-chan int) {
				defer wg.Done()
				for num := range ch {
					merged <- num
				}
			}(ch)
		}
		wg.Wait()
		close(merged)
	}()
	return merged
}

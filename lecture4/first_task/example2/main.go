package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func print(str string) {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println(str)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		print("JustCode1")
	}()
	go func() {
		defer wg.Done()
		print("JustCode2")
	}()
	go func() {
		defer wg.Done()
		print("JustCode3")
	}()
	wg.Wait()
	fmt.Println("Done!!!")
}

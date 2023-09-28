package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func print(str string) {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println(str)
	time.Sleep(time.Second)
	fmt.Println(str, "is printed")
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

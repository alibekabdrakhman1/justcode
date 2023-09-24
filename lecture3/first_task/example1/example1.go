package main

import "fmt"

func main() {
	ch := make(chan int)
	ch1 := make(chan string)
	go ch3(ch)
	go ch2(ch1)
	select {
	case r := <-ch1:
		fmt.Println(r)
	case r1 := <-ch:
		fmt.Println(r1)
	}

}
func ch3(ch chan int) {

}
func ch2(ch chan string) {}

package main

import "fmt"

func main() {
	i := 1
	ch := make(chan int)
	ch <- i
	ch <- i + 1
	fmt.Print(ch)
}

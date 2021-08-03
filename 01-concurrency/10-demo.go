package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	count := 10
	go fibonacci(ch, count)
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Exiting from main")
}

func fibonacci(ch chan int, count int) {
	x, y := 0, 1
	for idx := 0; idx < count; idx++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

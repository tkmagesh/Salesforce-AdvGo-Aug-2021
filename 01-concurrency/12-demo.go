package main

import (
	"fmt"
	"time"
)

func fibonacci(ch chan int) {
	x, y := 0, 1
	stop := time.After(10 * time.Second)
	for {
		select {
		case ch <- x:
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		case <-stop:
			fmt.Println("Quitting")
			close(ch)
			return
		}
	}
}

func main() {
	fmt.Println("Press ENTER to stop...")
	ch := make(chan int)
	go fibonacci(ch)
	go func() {
		for no := range ch {
			fmt.Println(no)
		}
	}()
	var input string
	fmt.Scanln(&input)
}

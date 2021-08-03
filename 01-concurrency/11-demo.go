package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	stop := make(chan bool)

	go fibonacci(ch, stop)
	go func() {
		for no := range ch {
			fmt.Println(no)
		}
	}()
	fmt.Println("press ENTER to stop...")
	var input string
	fmt.Scanln(&input)
	stop <- true
	fmt.Println("Exiting from main")
}

func fibonacci(ch chan int, stop chan bool) {

	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
			time.Sleep(500 * time.Millisecond)
		case <-stop:
			fmt.Println("Quitting")
			close(ch)
			return
		}
	}
}

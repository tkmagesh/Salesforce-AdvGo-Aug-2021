package main

import (
	"fmt"
	"time"
)

func main() {

	resultCh := make(chan int, 1)

	go add(10, 20, resultCh)
	//reading data from the channel
	time.Sleep(time.Second * 5)
	result := <-resultCh

	fmt.Println(result)
}

func add(x, y int, resultCh chan int) {
	result := x + y

	//writing data into the channel
	fmt.Println("Attempting to write data in the channel")
	resultCh <- result
	fmt.Println("Attemp to write data in the channel successful")

}

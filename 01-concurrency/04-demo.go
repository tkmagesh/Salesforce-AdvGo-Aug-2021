package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	resultCh := make(chan int)
	wg.Add(1)
	go add(10, 20, resultCh, wg)
	//reading data from the channel
	result := <-resultCh
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int, resultCh chan int, wg *sync.WaitGroup) {
	result := x + y

	//writing data into the channel
	resultCh <- result
	wg.Done()
}

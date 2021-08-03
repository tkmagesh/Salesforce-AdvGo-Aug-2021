package main

import (
	"fmt"
	"sync"
)

//commuicate by sharing memory
var result int
var mutex *sync.Mutex = &sync.Mutex{}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg)

	/* other operation */
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup) {
	mutex.Lock()
	result = x + y
	mutex.Unlock()
	wg.Done()
}

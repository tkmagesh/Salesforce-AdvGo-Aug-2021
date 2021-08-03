package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go print("Hello", wg, time.Second*5)

	wg.Add(1)
	go print("World", wg, time.Second*3)

	wg.Wait()
	fmt.Println("Exiting from main!")
}

func print(s string, wg *sync.WaitGroup, d time.Duration) {
	time.Sleep(d)
	fmt.Println(s)
	wg.Done()
}

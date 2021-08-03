package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	stop := time.After(20 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Print("Tick")
		case <-stop:
			fmt.Print("Done")
			return
		default:
			fmt.Print(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

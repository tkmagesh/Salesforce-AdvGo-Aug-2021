package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	go fn(ch)
	fmt.Println("Attempting to read 10")
	fmt.Println(<-ch)
	fmt.Println("Finished to read 10")
	fmt.Println("Attempting to read 20")
	fmt.Println(<-ch)
	fmt.Println("Finished to read 20")
	fmt.Println("Attempting to read 30")
	fmt.Println(<-ch)
	fmt.Println("Finished to read 30")
	fmt.Println("Attempting to read 40")
	fmt.Println(<-ch)
	fmt.Println("Finished to read 40")
	fmt.Println("Attempting to read 50")
	fmt.Println(<-ch)
	fmt.Println("Finished to read 50")
	fmt.Println("Attempting to read 60")
	fmt.Println(<-ch)
	fmt.Println("Finished to read 60")

}

func fn(ch chan int) {
	fmt.Println("Attempting to write 10")
	ch <- 10
	fmt.Println("Finished writing 10")
	fmt.Println("Attempting to write 20")
	ch <- 20
	fmt.Println("Finished writing 20")
	fmt.Println("Attempting to write 30")
	ch <- 30
	fmt.Println("Finished writing 30")
	fmt.Println("Attempting to write 40")
	ch <- 40
	fmt.Println("Finished writing 40")
	fmt.Println("Attempting to write 50")
	ch <- 50
	fmt.Println("Finished writing 50")
	fmt.Println("Attempting to write 60")
	ch <- 60
	fmt.Println("Finished writing 60")

}

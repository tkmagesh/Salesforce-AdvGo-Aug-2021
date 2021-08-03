/* divide & conquer using go routines*/
package main

import (
	"fmt"
)

func main() {
	resultCh1 := make(chan int)
	resultCh2 := make(chan int)
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}
	firstSet := data[:len(data)/2]
	secondSet := data[len(data)/2:]
	go sum(resultCh1, firstSet...)
	go sum(resultCh2, secondSet...)
	finalResult := <-resultCh1 + <-resultCh2
	fmt.Println(finalResult)
}

func sum(resultCh chan int, nos ...int) {
	result := 0
	for _, v := range nos {
		result += v
	}
	resultCh <- result
}

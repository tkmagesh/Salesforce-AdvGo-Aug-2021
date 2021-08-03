/* divide & conquer using go routines*/
package main

import "fmt"

func main() {
	resultCh := make(chan int)
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}
	firstSet := data[:len(data)/2]
	secondSet := data[len(data)/2:]
	go sum(resultCh, firstSet...)
	go sum(resultCh, secondSet...)
	finalResult := <-resultCh + <-resultCh
	fmt.Println(finalResult)
}

func sum(resultCh chan int, nos ...int) {
	result := 0
	for _, v := range nos {
		result += v
	}
	resultCh <- result
}

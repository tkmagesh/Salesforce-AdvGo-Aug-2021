package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	dataCh := make(chan int)
	fileWg := &sync.WaitGroup{}
	fileWg.Add(2)
	go source("data1.dat", dataCh, fileWg)
	go source("data2.dat", dataCh, fileWg)

	evenCh := make(chan int)
	oddCh := make(chan int)
	evenResultCh := make(chan int)
	oddResultCh := make(chan int)
	processWg := &sync.WaitGroup{}
	processWg.Add(4)
	go splitter(dataCh, evenCh, oddCh, processWg)
	go sum(evenCh, evenResultCh, processWg)
	go sum(oddCh, oddResultCh, processWg)
	go merge(evenResultCh, oddResultCh, "result.dat", processWg)

	fileWg.Wait()
	close(dataCh)
	processWg.Wait()
	fmt.Println("Job Done!")
}

func source(fileName string, ch chan int, wg *sync.WaitGroup) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		file.Close()
		wg.Done()
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		val, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		ch <- val
	}

}

func splitter(dataCh chan int, evenCh chan int, oddCh chan int, wg *sync.WaitGroup) {
	defer func() {
		close(evenCh)
		close(oddCh)
		wg.Done()
	}()
	for val := range dataCh {
		if val%2 == 0 {
			evenCh <- val
		} else {
			oddCh <- val
		}
	}
}

func sum(ch chan int, resultCh chan int, wg *sync.WaitGroup) {
	result := 0
	for no := range ch {
		result += no
	}
	resultCh <- result
	wg.Done()
}

func merge(evenResultCh chan int, oddResultCh chan int, resultFile string, wg *sync.WaitGroup) {
	resultFileHandle, resultFileErr := os.Create(resultFile)
	if resultFileErr != nil {
		panic(resultFileErr)
	}
	defer func() {
		resultFileHandle.Close()
		wg.Done()
	}()
	for i := 0; i < 2; i++ {
		select {
		case evenResult := <-evenResultCh:
			resultFileHandle.WriteString(fmt.Sprintf("Even Total : %v\n", evenResult))
		case oddResult := <-oddResultCh:
			resultFileHandle.WriteString(fmt.Sprintf("Odd Total : %v\n", oddResult))
		}
	}
}

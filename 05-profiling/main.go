package main

import "fmt"

func main() {
	nos := GenerateNos()
	fmt.Println(len(nos))
}

func GenerateNos() []int {
	var nos []int
	for i := 0; i < 10000; i++ {
		nos = append(nos, i)
	}
	return nos
}

func GenerateNosV2() []int {
	nos := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		nos = append(nos, i)
	}
	return nos
}

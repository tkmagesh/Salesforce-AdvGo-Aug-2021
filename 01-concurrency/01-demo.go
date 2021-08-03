package main

import (
	"fmt"
)

func main() {
	go print("Hello")
	go print("World")
	//time.Sleep(time.Second * 2)
	var input string
	fmt.Scanln(&input)
}

func print(s string) {
	fmt.Println(s)
}

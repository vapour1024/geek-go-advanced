package main

import "fmt"

func main() {
	var a string
	for i := 0; i < 500000; i++ {
		a += "1"
	}
	fmt.Println(a)
}

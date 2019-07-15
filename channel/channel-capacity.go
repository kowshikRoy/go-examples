package main

import "fmt"

func main() {
	intStream := make(chan int, 4)
	for i := 0; i <= 4; i++ {
		intStream <- i
	}
	fmt.Println("Ended\n")
}

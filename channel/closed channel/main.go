// https://play.golang.org/p/7NOEDMuiuep

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan int)
	ticker := time.NewTicker(time.Second)
	go func() {
		for i := 0; i < 5; i++ {
			<-ticker.C
			fmt.Printf("Pushing to Channel: %d\n", i)
			c <- i
		}
		fmt.Println("Closing Channel")
		close(c)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Poping from channel: ", <-c)
		}
		wg.Done()
	}()
	wg.Wait()

}

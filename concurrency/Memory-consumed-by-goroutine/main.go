/*
This program is used to calculate the memory allocated by per goroutines.
Link: https://play.golang.org/p/06QIWCFACEz
We use a channel to block all the goroutines.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}
	wg := sync.WaitGroup{}
	var c <-chan interface{}
	noop := func() { wg.Done(); <-c }

	const numOfGoRoutine = 1e5
	wg.Add(numOfGoRoutine)
	before := memConsumed()
	for i := 0; i < numOfGoRoutine; i++ {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb\n", float64(after-before)/numOfGoRoutine/1024)
}

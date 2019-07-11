/*
A deadlocked program is one in which all concurrent processes are waiting on one another.
https://play.golang.org/p/jX5SkjnPRDi
The first goroutine will lock a, then sleep for a second and try to lock b,
The second goroutine will lock b first, then sleep for a second, and try to lock a
There is no guarantee that deadlock will happen in every case, but a sleep of second,
Another goroutine will most likely to be executed, and lock another part.

Here is the code: https://play.golang.org/p/qzI0Ubv2zio
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type Number struct {
	mu  sync.Mutex
	val int
}

func main() {
	var wg sync.WaitGroup
	printSum := func(a, b *Number) {
		defer wg.Done()
		a.mu.Lock()
		defer a.mu.Unlock()

		time.Sleep(time.Second)

		b.mu.Lock()
		defer b.mu.Unlock()

		fmt.Printf("Sum is %d\n", a.val+b.val)
	}
	wg.Add(2)
	var a, b Number
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}


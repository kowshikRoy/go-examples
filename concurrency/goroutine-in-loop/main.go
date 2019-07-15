/*
goroutine is running a closure that has closed over the iteration variable str,
which has a type of string. As our loop iterates, salutation is being assigned to the
next string value in the slice literal. Because the goroutines being scheduled may run
at any point in time in the future, it is undetermined what values will be printed from
within the goroutine.
On my machine, there is a high probability the loop will exit
Goroutines before the goroutines are begun. This means the salutation variable falls out of
scope. What happens then? Can the goroutines still reference something that has
fallen out of scope? Won’t the goroutines be accessing memory that has potentially
been garbage collected?
This is an interesting side note about how Go manages memory. The Go runtime is
observant enough to know that a reference to the salutation variable is still being
held, and therefore will transfer the memory to the heap so that the goroutines can
continue to access it.

Usually on my machine, the loop exits before any goroutines begin running, so
salutation is transferred to the heap holding a reference to the last value in my
string slice, “good day.” And so I usually see “good day” printed three times. The
proper way to write this loop is to pass a copy of salutation into the closure so that
by the time the goroutine is run, it will be operating on the data from its iteration of
the loop:
*/
package main

import (
	"fmt"
	"sync"
)

func wrontExample() {
	wg := sync.WaitGroup{}
	for _, salutation := range []string{"California", "Seattle", "NY", "Washington"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Here we reference the loop variable salutation created by ranging over a string
			// slice.
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

func correctExample() {
	wg := sync.WaitGroup{}
	for _, salutation := range []string{"California", "Seattle", "NY", "Washington"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			// Here we reference the loop variable salutation created by ranging over a string
			// slice.
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
func main() {
	wrontExample()
	correctExample()
}


Goroutines has the reference of the variable that is used inside of
the goroutine. Any modification will be reflected in the programs
after that.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	str := "Hello"
	go func() {
		defer wg.Done()
		str = "Welcome"
	}()
	wg.Add(1)
	wg.Wait()
	fmt.Println(str)
}

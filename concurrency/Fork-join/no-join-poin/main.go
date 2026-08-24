package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	now := time.Now()

	wg.Add(1)
	go func() {
		defer wg.Done() // we know defer right it's finally for js in go
		work()
	}()

	wg.Wait()

	fmt.Println(time.Since(now))

}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("inside work")
}

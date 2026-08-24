package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
	}()

}

func task1() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("inside task 1")
}

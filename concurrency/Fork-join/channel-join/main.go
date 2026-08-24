package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	done := make(chan struct{})

	go func() {

		work()
		done <- struct{}{}
	}()

	<-done

	fmt.Println(time.Since(now))

}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("inside work")
}

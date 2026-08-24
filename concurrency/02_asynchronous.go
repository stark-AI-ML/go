package main

import (
	"fmt"
	"time"
)

func t1() {
	time.Sleep(1000 * time.Millisecond)

	fmt.Println("task 1")
}

func t2() {
	time.Sleep(2000 * time.Millisecond)

	fmt.Println("task 2")
}

func t3() {
	time.Sleep(500 * time.Millisecond)

	fmt.Println("task 3")
}

func t4() {
	time.Sleep(1000 * time.Microsecond)

	fmt.Println("task 4")
}

func main() {

	now := time.Now()

	go t1()
	go t2()
	go t3()
	go t4()

	time.Sleep(2001 * time.Millisecond)

	fmt.Println(time.Since(now))

}

// 14.778µs
// asynchronous output : not printing any output? at last why? did it even run the program?
//  without time.sleep() at last

// 2 . -----------------------------------------------------------
// task 4
// task 3
// task 1
// 1.000447892s

// time.sleep(1000ms) at last

// 3 . ----------------------------------------------------------
// task 4
// task 3
// task 1
// task 2
// 3.000199115s

// time.sleep(3000ms) at last

// main() doesn't wait for goroutines automatically.
// If main() exits before the goroutines finish, the entire program exits and unfinished goroutines are terminated.

// And:

// The goroutines are concurrent, but main needs some synchronization mechanism to know when they are finished.

// Using time.Sleep() is just a crude way of keeping main alive. It is not the proper solution.

package main

import (
	"fmt"
	"time"
)

func task1() {
	time.Sleep(1000 * time.Millisecond)

	fmt.Println("task 1")
}

func task2() {
	time.Sleep(2000 * time.Millisecond)

	fmt.Println("task 2")
}

func task3() {
	time.Sleep(500 * time.Millisecond)

	fmt.Println("task 3")
}

func task4() {
	time.Sleep(1000 * time.Microsecond)

	fmt.Println("task 4")
}

func main() {

	now := time.Now()

	task1()
	task2()
	task3()
	task4()

	fmt.Println(time.Since(now))

}

// task 1
// task 2
// task 3
// task 4
// 3.502830342s

// synchronous output

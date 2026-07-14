package main

import "fmt"

func main() {
	//declaration

	var a [3]string

	a[0] = "hello"
	a[1] = "rudresh"
	a[2] = "singh"

	fmt.Println(a[0], a[1], a[2])

	nums := [6]int{2, 3, 4, 5, 6, 6}

	// no need to put integer before as we are assinging the int to each index
	// var primes int= [3]int{2, 3, 5}

	var primes = [3]int{2, 3, 5}

	fmt.Print(nums, primes)
}

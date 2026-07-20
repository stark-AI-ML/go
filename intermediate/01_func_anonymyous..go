package main

import "fmt"

// what does this some represents ---> var sum func(x int, y int) int
// this is a declaration of a fx as a varialble and it expect the fx which holds same param same return type

func sqaureOfSum(
	sum func(x int, y int) int) int {
	return sum(5, 5) * sum(5, 5)
}

func main() {

	// var sum int = sum(x,y, int) int
	// {
	// 	return x+y
	// }

	sum := func(x int, y int) int { return x + y }

	fmt.Print(sqaureOfSum(sum))
}

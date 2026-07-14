package main

import "fmt"

//  well classic for go and quite a impresseive thing
// go can return multiple value

func swap(x, y string) (string, string) {
	return y, x
}

// reutrn experiment although as per docs use it on short fx an i can see why

func split(sum int) (x, y int) {

	x = sum / 2
	y = sum - x
	return // well it will return both x and y

	// notice we don't need to intialise x, y,
	// (it is taken from the return type x and y, which is already intialise with zero before hand like js do for thier fx (not talking about return type yk what I meant))

	// why not use it? : you have to deal with if there is some issue and your fx is long

}

func main() {

	//	a string, b string = swap("hii", "hello");

	// this := declaration shortucut for go it's same as the above, it takes it types based on right side of assignement data during compilation

	a, b := swap("hii", "hello")

	fmt.Println(a, b)

}

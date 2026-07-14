package main

import "fmt"

func main() {

	// slices mosty used with array or derived ds like strings and all

	/*
	 * slice doen't create new contaienr with same value instead
	 * it creates the pointer container which pointes to the array of that value

	 */
	nums := [6]int{2, 3, 4, 5, 6, 6}

	var s []int = nums[1:4]

	fmt.Println(nums)
	fmt.Println(s)

	// using slice ptr to change the value of the array

	s[1] = 10000

	fmt.Println(nums)
	fmt.Println(s)

	// [2 3 4 5 6 6]
	// [3 4 5]
	// [2 3 10000 5 6 6]
	// [3 10000 5]

}

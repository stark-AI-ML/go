package main

import "fmt"

// go doesn't have while and do while loop

func main() {

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}

	// while loop in for
	sum_w := 1

	for sum_w < 100 {
		sum_w += sum_w
		fmt.Println(sum_w)
	}


	// infinite loop for without a condition is an infinite loop
	// for{

	// }



	// similar to do while 

	sum_w = 0

	for{
		sum_w += sum_w
		if(sum_w < 100){
			break; 
		}
	}

}

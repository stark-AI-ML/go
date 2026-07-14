package main

import (
	"fmt"
	"runtime"
)

func main() {

	var sum int = 15

	if sum < 10 {
		fmt.Println(sum)
	}

	// here you can intialise the var inside the
	// the if statement
	if x := 10; x < sum {
		fmt.Println(x)
	}

	// Switch statement :

	fmt.Print("go run")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Print("os x")
		break
	case "linux":
		fmt.Print("Linux")
		break
	default:
		fmt.Printf("%s \n", os)
	} 

}

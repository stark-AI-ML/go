package main 

import "fmt"


func printN(x int) (y int) {

	if(x==0) return x

	fmt.Println(x) 

	return printN(x-1)
}


func main(){
	printN(10)
}

package main

import "fmt"

/* as of now i don't see any diffrence b/w this and c++ for pointer idea it's completely same

 */

func main() {

	x := 10

	p := &x

	fmt.Printf("value  : %v  address :  %v", *p, p)

}

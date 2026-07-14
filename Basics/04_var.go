package main

import "fmt"

// intialisation of the go is always zero state of that dataype
/**

* string : ""
* int : 0
* float :0.0

* bool : false\
 */

// this is the only way to assign value at global (i.e global var) in go
var a, b, c bool

var (
	x int     = 10322
	y bool    = false
	z int8    = 34 // an alias for byte
	f float32 = 12.32
)

func main() {

	i := 10 // type inference by go it will assign value and it can't be changes later
	fmt.Println(a, b, c, i, "\n")

	fmt.Printf("type of %T value %d \n", x, x)
	fmt.Printf("type of %T value %v \n", x, x) // see both %v and %d is working

	fmt.Printf("type of %T value %v \n", y, y) // %v is universal
	fmt.Printf("type of %T value %v \n", z, z)

	// type constants :

	const pi = 3.14  // can't be changed later evaluted during compile time so must assign the value

	// boolean, string, number  --- these are the only three value that constant can have
	fmt.Printf("type of %T value %v", pi)

}

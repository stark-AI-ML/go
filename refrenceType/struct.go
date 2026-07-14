package main

import "fmt"

type rectangle struct {
	length  int
	breadth int
}

// same as just syntax differ

var (
	vl = rectangle{1, 2}
	v2 = rectangle{length: 1}
	v3 = rectangle{}
	p  = &rectangle{1, 2}
)

func main() {

	fmt.Println(rectangle{3, 4})

	rect := rectangle{5, 8}
	rect.breadth = 10       // can change the value
	fmt.Print(rect.breadth) // can access the val

	fmt.Println(vl, p, v2, v3)

}

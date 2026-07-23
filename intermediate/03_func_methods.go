package main

import "fmt"

type Rectangle struct {
	X int
	Y int
}

func (area Rectangle) Area() int {
	return area.X * area.Y
}

func (area *Rectangle) change(len int, breadth int)
{
	area.X = len
	area.Y = breadth 
}

func main() {

	r1 := Rectangle{3, 2}

	fmt.Print(r1.Area())
}

package main

import "fmt"

type Rectangle struct {
	X int
	Y int
}


// declaration 

// func (receiver ReceiverType) MethodName(parameters) returnType {
//     // body
// }


func (area Rectangle) Area() int {
	return area.X * area.Y
}

func (area *Rectangle) change(len int, breadth int)
{
	area.X = len
	area.Y = breadth 
}


/*
--------------------------------------------------------------------------------------------------------------------
 although i have used both at a same time for learning puropose don't try to use both normal method and pointer 
 method in same code base don't do that

 when to use normal : 
	- well logical anwser is when you don't have to just read or acess the data no mutation

 when to use pointer method : 
	- well 
	- 1: mutation 			
	- 2: when your struct data is way too much long as we know normal makes a copy of data the reason why change don't
		effect the org value...so when it is large creating it as normal method and using it multiple times 
		will create multiple copy instances of it hence extra mem per instance you will create 
		
	- 3: when you have used pointer use pointer all along don't use both at ranodm 
----------------------------------------------------------------------------------------------------------------------
*/

func main() {

	r1 := Rectangle{3, 2}

	fmt.Print(r1.Area())
}

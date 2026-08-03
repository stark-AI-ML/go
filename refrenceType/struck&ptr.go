package main

import "fmt"


// caps represtents if variale even Vertext is exportable or not same for X and Y 
type Vertex struct {
	X int
	Y int
}

func main() {

	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)

}

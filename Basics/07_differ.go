package main

import "fmt"

func main() {

	// all defer reminds me of finally from node.js it's main functionality is this i  think but i also think
	// it's better approach to have fxs like this

	/* working is simple all differ fx goes to the stack mem, where it waits all other fxs to finish then during return
	no matte what it will run either program breaks, or returns successful

	the reason why i said it finally :-> db break ---- defer can release lock at last
	and we don't have to remember to put it on finally at last block just put after opening the block and then just close it

	*/

	defer fmt.Print("i should be first?, nah i am the last to execute")

	for i := 0; i < 10; i++ {
		defer fmt.Printf("i : %d \n", i)
	}
}

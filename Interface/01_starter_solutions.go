package main

import "fmt"

// now as we have seen the problem like we can't scale with diffrent increasing
// diffrent general method for each (well if it's for each it's not general anyway it's specfic: )
// hence no :-------- scaling

//  and if you make this general using datatype any it will take any int, string not the struct we have defined so no good

//  so here comes : our interface : ----------------------------------

/*
The Better Solution — Interface

Ask yourself:

		What does Notify() actually need?

		Does it need an Email?

		No.

		Does it need an SMS?

		No.

	It only needs something that can Send.

So define that behavior:
*/
type Email struct{}

func (Email) Send(message string) {
	fmt.Println("Email:", message)
}

package main

import "fmt"

// must Read
// The interface defines a requirement (a contract).
// The concrete type provides the implementation.
// An interface value remembers which concrete type it currently holds.
// When you call a method on the interface, Go forwards the call to that concrete type's method.

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

type Notifier interface {
	Send(message string)
}

// msg type 1:
type Email struct{}

func (Email) Send(message string) {
	fmt.Println("Email:", message)
}

// msg type 2:
type SMS struct{}

func (SMS) Send(message string) {
	fmt.Println("SMS:", message)
}

func Notify(n Notifier, msg string) {
	n.Send(msg)
}

func main() {

	email := Email{}
	sms := SMS{}

	Notify(email, "Welcome!")

	Notify(sms, "OTP: 1234")
}

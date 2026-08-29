package main

import "fmt"

type studyStatus int

/*
iota is a special Go keyword that starts at 0 and increments by 1 for each constant in the const block.

	So this becomes:

	Placed     = 0
	Confirmed  = 1
	Shipped    = 2
	Delivered  = 3
	Cancelled  = 4

	It's basically a convenient way to avoid manually writing:

	const (

		Placed    OrderStatus = 0
		Confirmed OrderStatus = 1
		Shipped   OrderStatus = 2
		Delivered OrderStatus = 3
		Cancelled OrderStatus = 4

	)
*/

const (
	Started studyStatus = iota // it's like indexing for the consts
	Pending
	Active
	Finished
)

func (s studyStatus) status_String() string {

	switch s {
	case Started:
		return "Started"
	case Pending:
		return "Pending"
	case Active:
		return "Active"
	case Finished:
		return "Finished"
	default:
		return "Unkown"
	}
}

func main() {

	status := Started

	if status == Started {
		fmt.Println("Your project has been started")
	}

}

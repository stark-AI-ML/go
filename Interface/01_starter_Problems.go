package main

import "fmt"

// --------------------
// PROBLEM SECTION
// --------------------

// Requirement 1:
// Initially, our application only supports Email notifications.

type email struct{}

func (email) Send(message string) {
	fmt.Println("Email:", message)
}

// notify() only accepts email.
func notify(e email, msg string) {
	e.Send(msg)
}

// ---------------------------------------------

// Requirement 2:
// Later, the application also needs SMS support.

type sms struct{}

func (sms) Send(message string) {
	fmt.Println("SMS:", message)
}

// Bad Solution #1:
//
// Create a separate function for every notification type.

func notifySMS(s sms, msg string) {
	s.Send(msg)
}

/*
Later more requirements come:

	- WhatsApp
	- Push Notification
	- Slack

Now the code becomes:

	notify()
	notifySMS()
	notifyWhatsApp()
	notifyPush()
	notifySlack()

Imagine supporting 15 notification types.
Creating a new notify function every time
doesn't scale.
*/

// ---------------------------------------------

// Bad Solution #2:
//
// Use any.
//
// func notify(n any, msg string) {}
//
// Now anything can be passed:
//
//	notify(10, "Hello")
//	notify("abc", "Hello")
//	notify([]int{}, "Hello")
//
// These are not notification services.
// Inside notify(), we cannot safely call Send().
// We lose compile-time type safety.

// ---------------------------------------------

func main() {

	e := email{}

	// Works because notify() accepts email.
	notify(e, "Welcome")

	s := sms{}

	// This DOES NOT work:
	//
	// notify(s, "Welcome")
	//
	// Error:
	// cannot use s (type sms) as type email in argument to notify

	// So we are forced to create another function.
	notifySMS(s, "Welcome")
}

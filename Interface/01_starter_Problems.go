package main

import "fmt"

//  first i will go with the problem :

// 1_prpblem_section  so here we have our 1: one requirement

// Email :

type email struct{}

func (email) Send(message string) {
	fmt.Println("Email:", message)
}

// 2_problem_section now  we have second 2: requirement

// SMS :
type sms struct{}

func (sms) Send(message string) {
	fmt.Println("SMS ", message)
}

// Bad Solution #1
func NotifySMS(s sms, msg string) {
	s.Send(msg)
}

/*
	Later someone adds:

		WhatsApp
		Push Notifications
		Slack

		Now your code becomes

		NotifyEmail()

		NotifySMS()

		NotifyWhatsApp()

		NotifySlack()

	Imagine having 15 notification types. This doesn't scale.
*/

// // Bad Solution #2

// func Notify(n any, msg string) {

// }

/*
	Now anything can be passed.

			Notify(10, "Hello")

			Notify("abc", "Hello")

			Notify([]int{}, "Hello")

	These aren't notification services, so inside Notify you can't safely call Send.

	You lose type safety.
*/

// common
func Notify(email Email, msg string) {
	email.Send(msg)
}

func main() {
	email := email{}
	Notify(Email(email), "Welcome")

	// but now let say you want to implement Notify for second requirement

	sms := sms{}

	// Notify(sms, "Welcome!")
	// error : cannot use sms (variable of struct type SMS) as Email value in argument to Notify
	/*
		Because
			Notify(email Email, msg string)
			only accepts Email.
	*/

	NotifySMS(sms, "Welcome")

}

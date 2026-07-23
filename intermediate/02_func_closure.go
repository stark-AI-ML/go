package main

import "fmt"

/*
func giftCard() func (int)int{

}

int func (int) int  ----- don't get confuse it is nothing but a return data type of the giftcard function :

so what it is : -> it is just a defination as we know func must have return datatype but when a fx return fx so
return data type becomes like this

	 func (int)int ----> is ! a declartion but return datatype defination where it says fallowing :

	 	- where func represent return type is fx  // but it didn't stop there it's strict for
		- params (int) -- represents param requires integer datatype // at last int
		- int  --- represents int return dataype of that fx that we are returning
*/

/* Clousure

see the idea is simple people just make it complex you will call the upper fx at one once as a simple priciple of program
fx will be intialsed and all the value will be created i.e var, here fx inside fx

so idea is simple you will have box

fx parent : [

	variable p
	fx  return variable - x
]

so in fx parent mem , --- each time you subract var p  with children fx,  mem of variable p is in parent

so it will be deducted itteratively and will not be reassigned with say 100 again instead

*****
100 - 10 = 90

next time it will be 90 only cuz childfx are inside parent mem and it is reinitialising not the parent fx

so until you create new parent fx there will be subraction from child to parent mem so it will be same as box contain both

*****

Each call for parent fx creates a NEW parent environment.

---------------------------------------------------------

Parent Environment #1

+-------------------------------------------+
| amount = 100                              |
|                                           |
| debitFunc --------------------------+      |
+-------------------------------------|-----+
                                      |
                                      | captures
                                      v
                              uses same amount

card1 -----> debitFunc

card1(10)
amount = 90

card1(20)
amount = 70

card1(30)
amount = 40

---------------------------------------------------------


Parent Environment #2

+-------------------------------------------+
| amount = 100                              |
|                                           |
| debitFunc --------------------------+      |
+-------------------------------------|-----+
                                      |
                                      | captures
                                      v
                              uses same amount

card2 -----> debitFunc

card2(5)
amount = 95

card2(50)
amount = 45

---------------------------------------------------------




*/

func giftCard() func(int) int {
	amount := 100

	debitFunc := func(debitAmount int) int {
		if amount >= debitAmount {
			amount -= debitAmount
			return amount
		}
		return 0
	}

	return debitFunc
}

func main() {
	card := giftCard()
	fmt.Println(card(10)) // 90
	fmt.Println(card(20)) // 70
	fmt.Println(card(50)) // 20
	fmt.Println(card(30)) // 0 (not enough balance)

}

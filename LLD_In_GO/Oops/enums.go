package main

import "fmt"

type OrderStatus int

/*

	Your const block
const (
    PLACED OrderStatus = iota
    CONFIRMED
    SHIPPED
    DELIVERED
    CANCELLED
)

This is a constant declaration block.

It's declaring multiple constants at once.

So conceptually:

PLACED     → constant
CONFIRMED  → constant
SHIPPED    → constant
DELIVERED  → constant
CANCELLED  → constant

Because of iota, their values become:

PLACED     = 0
CONFIRMED  = 1
SHIPPED    = 2
DELIVERED  = 3
CANCELLED  = 4
*/

const (
	PLACED OrderStatus = iota
	CONFIRMED
	SHIPPED
	DELIVERED
	CANCELLED
)

func (s OrderStatus) String() string {
	names := [...]string{"PLACED", "CONFIRMED", "SHIPPED", "DELIVERED", "CANCELLED"}
	if int(s) >= 0 && int(s) < len(names) {
		return names[s]
	}
	return "UNKNOWN"
}

type PaymentMethod struct {
	displayName string
	feePercent  float64
}

var (
	CREDIT_CARD = PaymentMethod{"Credit Card", 2.5}
	DEBIT_CARD  = PaymentMethod{"Debit Card", 1.0}
	UPI         = PaymentMethod{"UPI", 0.0}
	NET_BANKING = PaymentMethod{"Net Banking", 1.5}
)

func (p PaymentMethod) GetDisplayName() string { return p.displayName }
func (p PaymentMethod) GetFeePercent() float64 { return p.feePercent }

type Order struct {
	orderId       string
	status        OrderStatus
	paymentMethod PaymentMethod
	amount        float64
}

func NewOrder(orderId string, paymentMethod PaymentMethod, amount float64) *Order {
	return &Order{orderId: orderId, paymentMethod: paymentMethod, amount: amount, status: PLACED}
}

// just see the fx you get the smart way it handled itteration of order :
func (o *Order) AdvanceStatus() bool {
	switch o.status {
	case PLACED:
		o.status = CONFIRMED
		return true
	case CONFIRMED:
		o.status = SHIPPED
		return true
	case SHIPPED:
		o.status = DELIVERED
		return true
	default:
		return false
	}
}

func (o *Order) Cancel() bool {
	if o.status == PLACED || o.status == CONFIRMED {
		o.status = CANCELLED
		return true
	}
	return false
}

func (o *Order) GetTotalWithFees() float64 {
	return o.amount + (o.amount * o.paymentMethod.GetFeePercent() / 100)
}

func (o *Order) DisplayInfo() {
	fmt.Printf("Order %s | Status: %s | Payment: %s | Amount: $%.2f (with fees: $%.2f)\n",
		o.orderId, o.status, o.paymentMethod.GetDisplayName(), o.amount, o.GetTotalWithFees())
}

// Usage
func main() {
	order := NewOrder("ORD-001", CREDIT_CARD, 99.99)
	order.DisplayInfo()

	order.AdvanceStatus() // PLACED -> CONFIRMED
	order.AdvanceStatus() // CONFIRMED -> SHIPPED
	order.DisplayInfo()

	fmt.Println("Cancel after shipping:", order.Cancel()) // false
}

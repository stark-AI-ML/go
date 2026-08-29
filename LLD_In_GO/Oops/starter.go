package main

import (
	"fmt"
)

type FoodOrder struct {
	orderId      string
	customerName string
	items        []string
	totalAmount  float64
	isPlaced     bool
}

func NewFoodOrder(orderId, customerName string) *FoodOrder {
	return &FoodOrder{
		orderId:      orderId,
		customerName: customerName,
		items:        make([]string, 0),
		totalAmount:  0.0,
		isPlaced:     false,
	}
}

// Only allows adding items before the order is placed
func (f *FoodOrder) AddItem(name string, price float64) {
	if f.isPlaced {
		fmt.Println("Cannot modify a placed order.")
		return
	}
	f.items = append(f.items, name)
	f.totalAmount += price
}

// Places the order if it has at least one item and hasn't been placed yet
func (f *FoodOrder) PlaceOrder() bool {
	if f.isPlaced || len(f.items) == 0 {
		return false
	}
	f.isPlaced = true
	return true
}

func (f *FoodOrder) GetItemCount() int {
	return len(f.items)
}

func (f *FoodOrder) DisplayOrder() {
	status := "PENDING"
	if f.isPlaced {
		status = "PLACED"
	}
	fmt.Printf("Order %s (%s) - %s\n", f.orderId, f.customerName, status)
	for _, item := range f.items {
		fmt.Printf("  - %s\n", item)
	}
	fmt.Printf("  Total: $%.2f\n", f.totalAmount)
}

func main() {
	order1 := NewFoodOrder("ORD-101", "Alice")
	order1.AddItem("Pizza", 12.99)
	order1.AddItem("Garlic Bread", 4.99)
	order1.AddItem("Coke", 2.49)
	order1.PlaceOrder()

	order2 := NewFoodOrder("ORD-102", "Bob")
	order2.AddItem("Burger", 9.99)
	order2.AddItem("Fries", 3.99)

	order1.DisplayOrder()
	fmt.Println()
	order2.DisplayOrder()
}

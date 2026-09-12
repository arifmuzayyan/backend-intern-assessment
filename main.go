package main

import (
	"errors"
	"fmt"
)

type CartItem struct {
	ProductID string
	Name      string
	Price     float64
	Quantity  int
}

type Voucher struct {
	Code            string
	DiscountPercent float64
	MaxDiscount     float64
	MinPurchase     float64
}

func CalculateFinalPrice(
	items []CartItem,
	voucher *Voucher,
) (subtotal float64, discount float64, total float64, err error) {

	// Validate cart
	if len(items) == 0 {
		return 0, 0, 0, errors.New("cart cannot be empty")
	}

	// Calculate subtotal and validate items
	for _, item := range items {
		if item.Quantity <= 0 || item.Price < 0 {
			return 0, 0, 0, errors.New("invalid item price or quantity")
		}

		subtotal += item.Price * float64(item.Quantity)
	}

	// No voucher
	if voucher == nil {
		return subtotal, 0, subtotal, nil
	}

	// Voucher is not applicable
	if subtotal < voucher.MinPurchase {
		return subtotal, 0, subtotal, nil
	}

	// Calculate nominal discount
	discount = subtotal * (voucher.DiscountPercent / 100)

	// Apply maximum discount cap
	if discount > voucher.MaxDiscount {
		discount = voucher.MaxDiscount
	}

	// Prevent negative total
	if discount > subtotal {
		discount = subtotal
	}

	total = subtotal - discount

	return subtotal, discount, total, nil
}

func main() {
	// simple test case
	items := []CartItem{
		{
			ProductID: "P001",
			Name:      "Laptop",
			Price:     5_000_000,
			Quantity: 1,
		},
		{
			ProductID: "P002",
			Name:      "Mouse",
			Price:     250_000,
			Quantity: 2,
		},
	}

	voucher := &Voucher{
		Code:            "FLASH10",
		DiscountPercent: 10,
		MaxDiscount:     500_000,
		MinPurchase:     1_000_000,
	}

	subtotal, discount, total, err := CalculateFinalPrice(items, voucher)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Cart Summary:")
	fmt.Println("-------------")
	fmt.Printf("Items    : %d\n", len(items))
	fmt.Printf("Subtotal : Rp%.0f\n", subtotal)
	fmt.Printf("Discount : Rp%.0f\n", discount)
	fmt.Printf("Total    : Rp%.0f\n", total)
}

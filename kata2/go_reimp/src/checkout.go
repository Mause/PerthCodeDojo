package main

import (
	"errors"
	"fmt"
	"math"
)

type discounter func(map[string]int, float64) float64

type Checkout struct {
	cart               map[string]int
	prices             map[string]float64
	discounts          []float64
	discount_functions []discounter
}

func NewCheckout() *Checkout {
	c := &Checkout{
		cart:               make(map[string]int),
		discount_functions: make([]discounter, 0),
		discounts:          make([]float64, 0),
		prices: map[string]float64{
			"apple":  0.5,
			"cherry": 5,
			"mango":  3,
		},
	}
	c.prices["rotten_cherry"] = c.prices["cherry"] * .8

	return c
}

func (self *Checkout) AddToCart(item string, amount int) error {
	if self.prices[item] == 0 {
		return errors.New(fmt.Sprintf(
			"%s is an invalid item\n", item,
		))
	}

	self.cart[item] += amount

	return nil
}

func (self *Checkout) CalculateTotal() float64 {
	var total float64 = 0

	for key, value := range self.cart {
		total += float64(value) * self.prices[key]
	}

	for i := range self.discount_functions {
		total = self.discount_functions[i](
			self.cart, total,
		)
	}

	return total
}

func (self *Checkout) AddDiscountFunction(function discounter) {
	self.discount_functions = append(self.discount_functions, function)
}

func (self *Checkout) AddDiscount(percentage float64) {
	self.AddDiscountFunction(
		func(cart map[string]int, total float64) float64 {
			return total * percentage
		},
	)
}

func (self *Checkout) receipt() string {
	result := ""

	for name, amount := range self.cart {
		result += fmt.Sprintf(
			" * %d %ss for $%.2f",
			amount,
			name,
			self.prices[name],
		)
	}

	return result
}

func discount_apple(cart map[string]int, total float64) float64 {
	apples := cart["apple"]
	if apples == 0 {
		return total
	}

	return total - (math.Floor(float64(apples/3.0)) * 0.5)
}

func discount_cherry(cart map[string]int, total float64) float64 {
	cherries := cart["cherry"]
	if cherries == 0 {
		return total
	}

	return total - (math.Floor(float64(cherries/3.0)) * 7.5)
}

func main() {
	register := NewCheckout()

	register.AddDiscountFunction(discount_apple)
	register.AddDiscountFunction(discount_cherry)

	register.AddToCart("apple", 1)
	register.AddToCart("apple", 2)

	fmt.Printf(register.receipt() + "\n")
}

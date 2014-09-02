package main

import (
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
	return &Checkout{
		cart:               make(map[string]int),
		discount_functions: make([]discounter, 0),
		discounts:          make([]float64, 0),
		prices: map[string]float64{
			"apple":  0.5,
			"cherry": 5,
			"rotten_cherry": 5 * .8,
			"mango":  3,
		},
	}
}

func (self *Checkout) add_to_cart(item string, amount int) {
	if (self.prices[item] == 0) {
		fmt.Printf("%s is an invalid item\n", item)
		return
	}
	self.cart[item] += amount
}

func (self *Checkout) calculate_total() float64 {
	var total float64 = 0

	for key, value := range self.cart {
		total += float64(value) * self.prices[key]
	}

	for discount_function := range self.discount_functions {
		total = self.discount_functions[discount_function](
			self.cart, total,
		)
	}

	for discount := range self.discounts {
		total *= self.discounts[discount]
	}

	return total
}

func (self *Checkout) apply_discount_function(function discounter) {
	self.discount_functions = append(self.discount_functions, function)
}

func (self *Checkout) apply_discount(percentage float64) {
	self.discounts = append(self.discounts, percentage)
}

func (self *Checkout) receipt() (string) {
    result := "";

    for name, amount := range self.cart {
        result += fmt.Sprintf(
            " * %d %ss for $%.2f",
            amount,
            name,
            self.prices[name],
        )
    }

    return result;
}

func discount_apple(cart map[string]int, total float64) float64 {
	apples := cart["apple"]
	if apples == 0 { return total }

	return total - (math.Floor(float64(apples / 3)) * .5)
}

func discount_cherry(cart map[string]int, total float64) float64 {
	cherries := cart["cherry"]
	if cherries == 0 { return total }

	return total - math.Floor(float64(cherries/3.0)) * 7.5
}

func main() {
	register := NewCheckout()

	register.apply_discount_function(discount_apple)
	register.apply_discount_function(discount_cherry)

	register.add_to_cart("apple", 1)
	register.add_to_cart("apple", 2)

	fmt.Printf(
		"%f\n",
		register.calculate_total(),
	)
}

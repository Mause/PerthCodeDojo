package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertTotal(t *testing.T, register *Checkout, price float64) {
	calculated := register.calculate_total()

	assert.Equal(
		t,
		calculated,
		price,
		fmt.Sprintf(
			"calculated $%.2f != $%.2f",
			calculated,
			price,
		),
	)
}

func TestEmptyCart(t *testing.T) {
	register := NewCheckout()
	assertTotal(t, register, 0)
}

func TestBasicCheckout(t *testing.T) {
	register := NewCheckout()
	register.add_to_cart("apple", 1)
	assertTotal(t, register, .5)
}

func TestThirdFreeCheckout(t *testing.T) {
	register := NewCheckout()
	register.add_discount_function(discount_apple)

	register.add_to_cart("apple", 5)
	assertTotal(t, register, 4.0*0.5)
}

func TestSevenApples(t *testing.T) {
	register := NewCheckout()
	register.add_discount_function(discount_apple)

	register.add_to_cart("apple", 7)
	assertTotal(t, register, 2.5)
}

func TestMultipleFruit(t *testing.T) {
	register := NewCheckout()
	register.add_discount_function(discount_apple)
	register.add_discount_function(discount_cherry)

	register.add_to_cart("apple", 2)
	register.add_to_cart("mango", 4)
	register.add_to_cart("cherry", 10)
	assertTotal(
		t,
		register,

		50-(7.5*3)+ // cherries
			12+ // mangoes 3 * 4
			1, // apples 2 * 0.5
	)
}

func TestCheaperCherries(t *testing.T) {
	register := NewCheckout()
	register.add_discount_function(discount_cherry)

	register.add_to_cart("cherry", 3)
	assertTotal(t, register, 15-7.5)
}

func TestCheaperCherriesAssumption(t *testing.T) {
	register := NewCheckout()
	register.add_discount_function(discount_cherry)

	register.add_to_cart("cherry", 6)
	assertTotal(t, register, 15)
}

func TestVeryBasic(t *testing.T) {
	register := NewCheckout()
	register.add_discount_function(discount_apple)
	register.add_discount_function(discount_cherry)

	register.add_to_cart("apple", 2)
	register.add_to_cart("cherry", 3)
	assertTotal(t, register, 7.5+1)
}

func TestSplitApples(t *testing.T) {
	register := NewCheckout()
	register.add_discount_function(discount_apple)

	register.add_to_cart("apple", 2)
	register.add_to_cart("apple", 1)
	assertTotal(t, register, 2*.5)
}

func TestDiscountedCherry(t *testing.T) {
	register := NewCheckout()
	register.add_discount_function(discount_cherry)

	register.add_to_cart("rotten_cherry", 5)
	assertTotal(t, register, 25*.8)
}

func TestTotalDiscount(t *testing.T) {
	register := NewCheckout()
	register.add_discount_function(discount_apple)

	register.add_to_cart("cherry", 2)
	register.add_to_cart("apple", 6)
	register.add_to_cart("mango", 1)
	register.add_discount(.2)

	assertTotal(t, register, 3)
}

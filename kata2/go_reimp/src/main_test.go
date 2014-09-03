package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertTotal(t *testing.T, register *Checkout, price float64) {
	calculated := register.CalculateTotal()

	assert.Equal(
		t,

		calculated, price,

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
	register.AddToCart("apple", 1)
	assertTotal(t, register, .5)
}

func TestThirdFreeCheckout(t *testing.T) {
	register := NewCheckout()
	register.AddDiscountFunction(discount_apple)

	register.AddToCart("apple", 5)
	assertTotal(t, register, 4.0*0.5)
}

func TestSevenApples(t *testing.T) {
	register := NewCheckout()
	register.AddDiscountFunction(discount_apple)

	register.AddToCart("apple", 7)
	assertTotal(t, register, 2.5)
}

func TestMultipleFruit(t *testing.T) {
	register := NewCheckout()
	register.AddDiscountFunction(discount_apple)
	register.AddDiscountFunction(discount_cherry)

	register.AddToCart("apple", 2)
	register.AddToCart("mango", 4)
	register.AddToCart("cherry", 10)
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
	register.AddDiscountFunction(discount_cherry)

	register.AddToCart("cherry", 3)
	assertTotal(t, register, 15-7.5)
}

func TestCheaperCherriesAssumption(t *testing.T) {
	register := NewCheckout()
	register.AddDiscountFunction(discount_cherry)

	register.AddToCart("cherry", 6)
	assertTotal(t, register, 15)
}

func TestVeryBasic(t *testing.T) {
	register := NewCheckout()
	register.AddDiscountFunction(discount_apple)
	register.AddDiscountFunction(discount_cherry)

	register.AddToCart("apple", 2)
	register.AddToCart("cherry", 3)
	assertTotal(t, register, 7.5+1)
}

func TestSplitApples(t *testing.T) {
	register := NewCheckout()
	register.AddDiscountFunction(discount_apple)

	register.AddToCart("apple", 2)
	register.AddToCart("apple", 1)
	assertTotal(t, register, 2*0.5)
}

func TestDiscountedCherry(t *testing.T) {
	register := NewCheckout()
	register.AddDiscountFunction(discount_cherry)

	register.AddToCart("rotten_cherry", 5)
	assertTotal(t, register, 25*0.8)
}

func TestTotalDiscount(t *testing.T) {
	register := NewCheckout()
	register.AddDiscountFunction(discount_apple)

	register.AddToCart("cherry", 2)
	register.AddToCart("apple", 6)
	register.AddToCart("mango", 1)
	register.AddDiscount(.2)

	assertTotal(t, register, 3)
}

func TestErrorOnInvalidItem(t *testing.T) {
	register := NewCheckout()

	err := register.AddToCart("invalid", 0)

	assert.Error(t, err, "Shouldn't accept an item we don't recognize")
}

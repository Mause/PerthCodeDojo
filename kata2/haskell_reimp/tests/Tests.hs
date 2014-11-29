module Main where

import Checkout
import Test.Hspec

buildCheckout :: [(String,Int)] -> [DiscountFunction] -> Checkout
buildCheckout items funcs = do
    let register = newCheckout
        registerWithFuncs = addDiscountFunctions register funcs
    addManyToCart registerWithFuncs items


runTest :: [(String,Int)] -> [DiscountFunction] -> Double -> Expectation
runTest items funcs res = do
    calculateTotal (buildCheckout items funcs) `shouldBe` res


main :: IO ()
main = hspec $ do
    describe "Misc:" $ do
        it "An empty cart returns zero" $ do
            runTest [] [] 0

        it "test_very_basic" $ do
            runTest [("cherry", 3), ("apple", 2)] [discountApple, discountCherry] (7.5 + 1)

    describe "Apples:" $ do
        it "one apple should cost 50c" $ do
            runTest [("apple",1)] [] 0.5

        it "three apples, should have the third free" $ do
            runTest [("apple", 5)] [discountApple] (4 * 0.5)

        it "test_seven_apples" $ do
            runTest [("apple", 7)] [discountApple] 2.5

    describe "Multiple fruit:" $ do
        it "should equal enough... :P" $ do
            let items = [("apple", 2), ("mango", 4), ("cherry", 10)]
                funcs = [discountCherry, discountApple]
                result = (
                          50 - (7.5 * 3) + -- cherries
                          12 +             -- mangoes 3 * 4
                          1                -- apples 2 * 0.5
                         )
            runTest items funcs result

    describe "Discounts:" $ do
        it "test_cheaper_cherries" $ do
            runTest [("cherry", 3)] [discountCherry] (15 - 7.5)

        it "test_cheaper_cherries_assumption" $ do
            runTest [("cherry", 6)] [discountCherry] 15

        it "test_split_apples" $ do
            runTest [("apple", 2), ("apple", 1)] [discountApple] (2 * 0.5)

        it "five discounted cherries should cost 25 * .8" $ do
            runTest [("rotten_cherry", 5)] [discountCherry] (25 * 0.8)

        it "test_total_discount" $ do
            let reg = buildCheckout [("apple",2),("cherry",2),("mango",1)] [discountApple, discountCherry]
            calculateTotal (addDiscount reg 0.5) `shouldBe` 7

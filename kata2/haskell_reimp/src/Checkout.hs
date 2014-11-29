module Checkout where

import qualified Data.Map as DM

type DiscountFunction = Cart -> Double -> Double
type Cart             = DM.Map String Int
type Checkout         = (Cart, [DiscountFunction])

addToCart :: Checkout -> String -> Int -> Checkout
addToCart register name num = do
    let newCart = DM.insertWith (+) name num (getCart register) :: Cart
    (newCart, (getDiscountFunctions register))


addManyToCart :: Checkout -> [(String,Int)] -> Checkout
addManyToCart register [] = do register
addManyToCart register (item:items) = do
    addManyToCart (addToCart register (fst item) (snd item)) items


getPrice :: String -> Double
getPrice name
    | name == "apple"         = 0.5
    | name == "cherry"        = 5.0
    | name == "mango"         = 3.0
    | name == "rotten_cherry" = (getPrice "cherry") * 0.8
    | otherwise               = error "Invalid product"


summer :: String -> Int -> Maybe Double
summer key num = Just $ (getPrice key) * fromIntegral num


newCheckout :: Checkout
newCheckout = (DM.empty, [])


getCart :: Checkout -> Cart
getCart = fst


getDiscountFunctions :: Checkout -> [DiscountFunction]
getDiscountFunctions = snd


calcDiscount :: (DiscountFunction, Cart) -> Double -> Double
calcDiscount (func, cart) total = total - (func cart total)


calculateTotal :: Checkout -> Double
calculateTotal register = do
    let individual = DM.elems (DM.mapMaybeWithKey summer (getCart register))

        -- build a list of tuples with each function accompanied by it's own
        -- reference(?) to the cart
        funcsWithCart = zip (getDiscountFunctions register) (repeat (getCart register))

    foldr calcDiscount (sum individual) funcsWithCart


addDiscountFunction :: Checkout -> DiscountFunction -> Checkout
addDiscountFunction register func = do
    (getCart register, func : (getDiscountFunctions register))


addDiscountFunctions :: Checkout -> [DiscountFunction] -> Checkout
addDiscountFunctions register [] = register
addDiscountFunctions register (func:funcs) = do
    addDiscountFunctions (addDiscountFunction register func) funcs


addDiscount :: Checkout -> Double -> Checkout
addDiscount register percent = do
    addDiscountFunction register (\ _ total -> total - (total * percent))


quotInteg :: Int -> Double
quotInteg n = fromIntegral (n `quot` 3)


discount :: String -> Double -> DiscountFunction
discount key mul_val cart _ = do
    case DM.lookup key cart of
        Nothing -> 0
        Just val -> do quotInteg val * mul_val


-- the next two functions are partial applications of the above "discount"

-- every third apple is free
discountApple :: DiscountFunction
discountApple = discount "apple" 0.5


-- every third cherry is half price
discountCherry :: DiscountFunction
discountCherry = discount "cherry" 7.5

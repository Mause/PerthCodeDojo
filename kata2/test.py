import unittest
from checkout import Apple, Cherry, Checkout, Mango


class TestCheckout(unittest.TestCase):
    def setUp(self):
        self.register = Checkout()

    def assertTotal(self, price):
        self.assertEqual(self.register.calculate_total(), price)

    def test_empty_cart(self):
        self.assertTotal(0)

    def test_basic_checkout(self):
        self.register.add_to_cart(Apple(1))
        self.assertTotal(.5)

    def test_third_free_checkout(self):
        self.register.add_to_cart(Apple(5))
        self.assertTotal(4 * 0.5)

    def test_seven_apples(self):
        self.register.add_to_cart(Apple(7))
        self.assertTotal(2.5)

    def test_multiple_fruit(self):
        self.register.add_multiple([
            Apple(2),
            Mango(4),
            Cherry(10)
        ])
        self.assertTotal(
            50 - (7.5 * 3) +  # cherries
            12 +  # mangoes 3 * 4
            1     # apples 2 * 0.5
        )

    def test_cheaper_cherries(self):
        self.register.add_to_cart(Cherry(3))
        self.assertTotal(15 - 7.5)

    def test_cheaper_cherries_assumption(self):
        self.register.add_to_cart(Cherry(6))
        self.assertTotal(15)


def main():
    unittest.main()

if __name__ == '__main__':
    main()

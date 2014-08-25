import unittest
from checkout import Apple, Cherry, Checkout, Mango


class TestCheckout(unittest.TestCase):
    def setUp(self):
        self.register = Checkout()

    def assertTotal(self, price):
        self.assertEqual(self.register.calculate_total(), price)

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
            50 +  # cherries 5 * 10
            12 +  # mangoes 3 * 4
            1     # apples 2 * 0.5
        )


def main():
    unittest.main()

if __name__ == '__main__':
    main()

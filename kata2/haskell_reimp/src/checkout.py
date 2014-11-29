import math


class Checkout(object):
    def __init__(self):
        self.items = []
        self.discounts = []

    def add_to_cart(self, item):
        assert isinstance(item, Item)
        self.items.append(item)

    def add_multiple(self, items):
        self.items.extend(items)

    def calculate_total(self) -> int:
        total = sum(
            item.calc_price()
            for item in self.items
        )

        for discount in self.discounts:
            total *= discount

        return total

    def apply_discount(self, percentage):
        self.discounts.append(percentage)

    def reciept(self):
        return '\n'.join(
            ' * {} {}s for ${:.2f}'.format(
                item.num,
                item.name,
                item.calc_price()
            )
            for item in self.items
        )


class Item(object):
    def __init__(self, name, num, price):
        self.name = name
        self.num = num
        self.price = price

    def calc_price(self):
        return self.num * self.price


class Fruit(Item):
    pass


class Apple(Fruit):
    def __init__(self, num):
        super().__init__('Apple', num, 0.5)

    def calc_price(self):
        duos, extra = divmod(self.num, 3)

        return (duos * (2 * 0.5)) + (extra * .5)


class Cherry(Fruit):
    def __init__(self, num):
        super().__init__('Cherry', num, 5)

    def apply_discount(self, total):
        discount = math.floor(self.num / 3) * 7.5

        return total - discount

    def calc_price(self):
        total = super().calc_price()

        return self.apply_discount(total)


class RottenCherry(Cherry):
    def apply_discount(self, total):
        return total * .8


class Mango(Fruit):
    def __init__(self, num):
        super().__init__('Mango', num, 3)

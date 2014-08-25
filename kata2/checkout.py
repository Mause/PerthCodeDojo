

class Checkout(object):
    def __init__(self):
        self.items = []

    def add_to_cart(self, item):
        assert isinstance(item, Item)
        self.items.append(item)

    def calculate_total(self) -> int:
        return sum(
            item.calc_price()
            for item in self.items
        )


class Item(object):
    def __init__(self, name, num, price):
        self.name = name
        self.num = num
        self.price = price

    def calc_price(self):
        return self.num


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


class Mango(Fruit):
    def __init__(self, num):
        super().__init__('Mango', num, 3)

class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age

    def __str__(self):
        return f"{self.name} is {self.age} years old"

p1 = Person("John", 36)
print(p1)

def calculateTotalPrice(quantity:float, price:int)-> float:
    totalPrice = quantity * price
    return totalPrice
result = calculateTotalPrice(55.5, 100)

print(f"The total price is: {result}")
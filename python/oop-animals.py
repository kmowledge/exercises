from abc import ABC, abstractmethod

class Animal(ABC):
      @abstractmethod
      def sound(self):
            pass

class Dog(Animal):
      def sound(self):
            return "Hau!"
      def aport(self):
            print("<aportuje>")
      
class Cat(Animal):
      def sound(self):
            return "Meow?"
      def cuddle(self):
            print("<łasi się>")

dog1 = Dog()
cat1 = Cat()

print(cat1.sound())
print(dog1.sound())

#numerous inheritance
class Pet(Cat, Dog): #inherits methods after both, but if a method overlaps it takes the latest defined, which overwrites the previous one
      pass
      # self.name = name 
      # self = 

my_dog = Pet()
print(my_dog.sound())
my_dog.aport()
my_dog.cuddle()
package animals

import "fmt"

// Dog - This is a Dog! It's a struct type that represents a dog with a name, breed, and age.
type Dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
}

// NewDog is a factory function that creates a new Dog instance.
func NewDog(name, breed string, age int) Dog {
	return Dog{
		Name:  name,
		Breed: breed,
		Age:   age,
	}
}

// Bark is a method of the Dog type that returns a string about the dog barking.
// In Go, methods are functions that have a receiver type, which is the type that the method is associated with.
// The receiver type is specified before the function name, and it allows the method to access the fields of the Dog instance.
// To invoke this method, you would use a Dog instance followed by a dot and the method name.
func (d Dog) Bark() string {
	barkingMessage := fmt.Sprintf("%s says woof woof!", d.Name)
	return barkingMessage
}

// VVVV    BOTH OF THESE FUNCTIONS WOULD USUALLY BE IN A DIFFERENT FILE    VVVV

// StaticBark is a function that creates a barking string from some "name" input.
// "Static" is not a common Go idiom, but if you want a function that behaves like a static method,
// you can define it outside the Dog struct by excluding the receiver type.
func StaticBark(name string) string {
	if name == "" {
		return "Woof woof!"
	}
	return fmt.Sprintf("%s says woof woof!", name)
}

// StillStaticBark is still behaving like a static method, but it takes a Dog instance as an argument.
// This basically does the same thing as Bark, but it is not a method of the Dog type because
// it lacks a receiver type.
func StillStaticBark(d Dog) string {
	if d.Name == "" {
		return "Woof woof!"
	}
	return fmt.Sprintf("%s says woof woof!", d.Name)
}

package animals

import "fmt"

type Dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
}

func NewDog(name, breed string, age int) Dog {
	return Dog{
		Name:  name,
		Breed: breed,
		Age:   age,
	}
}

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

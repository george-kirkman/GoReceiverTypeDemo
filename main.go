package main

import "receiverTypeDemo/animals"

func main() {
	// Make a new Dog instance with the factory function
	dog := animals.NewDog("Fido", "Golden Retriever", 5)

	// You can access methods with the Dog receiver type by starting with the dog variable
	println(dog.Bark())

	// If you wanted to use "static" functions, use the package name
	println(animals.StaticBark("Poochie"))

	// I also defined a different "static" function that takes a Dog instance.
	// This basically does the same thing as Bark, but it is not a method of the Dog type
	// and is invoked in a different way, with the package name and the Dog instance as an argument.
	println(animals.StillStaticBark(dog))
}

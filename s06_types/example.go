package main

import "fmt"

type Location string

// this is called the special argument we have here.
// In a method, one argument is moved from the argument list to a special position before the function name, which connects the function to a specific type, making it a method.
// Use a noun that refers to the type the method is being added to, typically a lowercase identifier representing the type or origin.
// Methods are called using dot notation, such as 'nyc.DistanceTo(destination)', where the type instance precedes the method name.
func (location Location) DistanceTo(destination Location) distance {
	// calculation here ....
	fmt.Printf("Origin %v Destination %v\n", location, destination)
	return 10
}

func locationTest() {
	nyc := Location("34.99, 889.00")
	nyc.DistanceTo(Location("78.99, 89.99"))

}

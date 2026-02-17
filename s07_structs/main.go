package main

import (
	"fmt"
	"mod7/data"
)

// Structures are kind of the replacement of a class. It's just a data type with properties strongly typed. It's like you merge a couple of properties under one name.
// We'll see how we can add methods to structures because it's a type. (See last lecture)
// in fact, you don't need to pass all the properties.
// Go does not support traditional inheritance. Instead, it has a concept called embedding that can simulate some inheritance-like behaviors

func main() {
	max := data.Instructor{}
	max.FirstName = "Maximiliano"
	max.LastName = "Firtman"
	// some values can left as well if we want.

	// calling the nmethod over the struct
	fmt.Println(max.PrintInstructor())

	// another way of setting the values.
	rahul := data.Instructor{Id: 89, FirstName: "rahul"}

	fmt.Println(data.Instructor.PrintInstructor(rahul))

	fmt.Println(rahul)

	fmt.Println(max)
}

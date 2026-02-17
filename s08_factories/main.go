package data

import (
	"fmt"
	teacher "mod8/Example"
)

// we don't have the constructur in the struct but we have the design pattern to create the factories.
// It is method which will create the struct, we will use the method name starting with capital letter 'New' as the prefix. and this will return the struct.

type Instructor struct {
	Id        int
	FirstName string // IMPORTANT : note when we have export to another package every value inside the struct we need to write in capital first letter.
	LastName  string
	Score     int
}

// factory method : it is basically a method which returns the new struct.
// they make sure that the values are passed in the least required values are added in the struct.
func NewInstructor(firstName string, lastName string) Instructor {
	return Instructor{FirstName: firstName, LastName: lastName}
}

func main() {
	// here kyle is the new instructor.
	kyle := NewInstructor("Kyle", "Simpson")

	fmt.Println(kyle)

	// note : Some people, some developers are creating a data package in general.Some developers are creating one package for per model .So then the package can be called instructor. If that's the case, then you can use directly something like 'instructor.New' . So you don't need to specify the type of element, the structure you want, okay? So you can say 'structure.New' . By that case, you need a different package per model, per piece of information that you are saving in your app.

	// example here : (check the example folder for code)
	// here rahul is the new teacher and this package has the only struct of type teacher and a method New (factory method) which creates the teacher and it is capital so that method is exported.
	rahul := teacher.New("rahul", "kumar")
	fmt.Println(rahul)

}

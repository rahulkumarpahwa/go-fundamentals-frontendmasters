package main

import (
	"fmt"
	// "mod9/course"
	// "mod9/instructor"
	"mod9/instructor"
	"mod9/workshop"
	// "time"
)

func main() {
	// instructor := instructor.Instructor{FirstName: "Kyle", LastName: "Simpson"}

	// goCourse := course.Course{Name: "go fundamentals", Instructor: instructor}

	// GoWorkshopClass := workshop.Workshop{Course: goCourse, Date: time.Now()}

	// now here the workshop.Workshop struct does not have the stringer method which will call the string() method over this printing. so, the solution is to use the embedding.
	// fmt.Printf("%v", GoWorkshopClass)

	// new way of creating the workshop variable with embedding:
	// this Workshop struct now has the all the properties in the Course Struct available in it due to EMBEDDING. we can pass them directly without creating the Course and Intructor first.
	// now we can directly access the properties of the Course with Dot (.) operator and similar for the Intructor as well.
	// ALL the properties from the Course are now part of the Workshop.
	// note : don't try to pass 'Course' properties in 'workshop.Workshop{}' directly here, it will not work.
	GoWorkshopClass2 := workshop.Workshop{}
	GoWorkshopClass2.Name = "go fundamentals"
	GoWorkshopClass2.Instructor.FirstName = "Kyle"
	GoWorkshopClass2.Instructor.LastName = "Simspon"
	GoWorkshopClass2.Legacy = true
	fmt.Printf("%v", GoWorkshopClass2)

	jsInstructor := instructor.Instructor{FirstName: "Akshay", LastName: "Saini"}
	jsWorkShop := workshop.NewWorkShop("Js WorkShop", jsInstructor)
	fmt.Println(jsWorkShop) // here the stringer method of the Workshop will be called.
	// also, note that when we embedded type then we also embedded thier stringer methods as well.
}

// What is embedding in Go and how does it differ from traditional inheritance?
//1. Embedding in Go allows one type to include another type's properties and methods by directly including the type without an identifier. Unlike traditional inheritance, it provides a way to compose types by adding all properties and methods of the embedded type to the parent type.

//2. When embedding types with naming collisions, you need to disambiguate by explicitly referencing the specific embedded type's property, such as .instructor or .course, to resolve which property you want to use.

// 3. When embedding a type, its methods are also embedded and become available to the parent type, allowing the parent type to use methods of the embedded type directly.

// 4. Embedding allows code reuse without duplicating type definitions, provides a clean way to compose types, automatically includes properties and methods of the embedded type, and helps avoid code duplication.


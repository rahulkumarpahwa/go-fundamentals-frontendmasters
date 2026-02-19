package main

import (
	"fmt"
	"mod10/course"
	"mod10/instructor"
	"mod10/interfaces"
	"mod10/workshop"
)

func main() {

	goCourse := course.Course{Name: "go fundamentals", Instructor: instructor.Instructor{FirstName: "Kyle", LastName: "Simpson"}}

	jsWorkShop := workshop.Workshop{}
	jsWorkShop.Name = "js fundamentals"
	jsWorkShop.Instructor.FirstName = "Akshay"
	jsWorkShop.Instructor.LastName = "Saini"

	// Now, suppose we have to create the slice of the courses.
	// then jsWorkshop is not a course but goCourse is :
	var courses [2]course.Course
	courses[0] = goCourse
	// courses[1] = jsWorkShop // Type Mismatch here.

	// and we want to loop the courses to show:
	for _, course := range courses {
		fmt.Println(course)
	}

	// so, to print the jsWorkshop here we will use the interfaces here.

	var courses2 [2]interfaces.Signable
	courses2[0] = goCourse
	courses2[1] = jsWorkShop
	// now we can see that the both goCourse and jsWorkShop work under the same window, using interfaces which has the method SignUp and both needs to create that in them.

	// now we can loop over them to print:
	for _, courseAndWorkShop := range courses2 {
		fmt.Println(courseAndWorkShop)
	}

	// IMPLICIT IMPLEMENTATION:
	// Note: Where did I set that course I worshop are implementing that interface.I'm not saying it anywhere. "In Go, it's implicit by the name or the signature of your interface". So, if you have any type of with that method, with that signature you're in, you're in the interface automatically. You don't need to say so.You don't need to say implement 'Signable'.

	// So, we can have the  Collection of two different types of Struct using the interface and it creates the runtime polymorphism.
	// Looping through a lot of objects without actually knowing the kind of objects,I just need to know that it has the function or the action that I need.
}

package main

import (
	"fmt"
	"mod9/course"
	"mod9/instructor"
	"mod9/workshop"
	"time"
)

func main() {
	// instructor := instructor.Instructor{FirstName: "Kyle", LastName: "Simpson"}

	// goCourse := course.Course{Name: "go fundamentals", Instructor: instructor}

	// GoWorkshopClass := workshop.Workshop{Course: goCourse, Date: time.Now()}

	// now here the workshop.Workshop struct does not have the stringer method which will call the string() method over this printing. so, the solution is to use the embedding.
	// fmt.Printf("%v", GoWorkshopClass)

	// new way of creating the workshop variable with embedding:
	// this Workshop struct now has the all the properties in the Course Struct available in it due to EMBEDDING. we can pass them directly without creating the course first.
	GoWorkshopClass2 := workshop.Workshop{Name: "Go Course second", }
	fmt.Printf("%v", GoWorkshopClass2)

}

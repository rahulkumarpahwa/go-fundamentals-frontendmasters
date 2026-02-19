package workshop

import (
	"fmt"
	"mod9/course"
	"mod9/instructor"
	"time"
)

// here the workshop is same as the course in frontend masters but it is live. ie. it has the property date in it.

// type Workshop struct {
// 	Course course.Course
// 	Date time.Time
// }

// here we will instead of passing the Course as the property with type course.Course we will only put the type as. this is called Embedding :

type Workshop struct {
	course.Course // embedding
	Date          time.Time
}

// this embedding will make the properties / datatype available in the Workshop from the course with taking the property of the course.
// see working in the main.go

// now we will create here the Stringer() method for the Workshop struct as :
func (w Workshop) String() string {
	return fmt.Sprintf("\nName of the Course : %v\n Instructor of the Course : %v", w.Name, w.Instructor.FirstName+" "+w.Instructor.LastName)
}

// Now we will create the function which is called the factory method, which will create the new course structures.
func NewWorkShop(name string, instructor instructor.Instructor) Workshop {
	w := Workshop{}
	w.Name = name
	w.Instructor = instructor
	return w
}
// check main.go for its usage.

// note: In case of multiple same properties in the struct itself and the struct one which is embedded then we can use the '.' dot operator to get things unambigious, to access the property of the particular struct.

package workshop

import (
	"fmt"
	"mod10/course"
	"mod10/instructor"
	"time"
)

type Workshop struct {
	course.Course
	Date time.Time
}

func (w Workshop) String() string {
	return fmt.Sprintf("\nName of the Course : %v\n Instructor of the Course : %v", w.Name, w.Instructor.FirstName+" "+w.Instructor.LastName)
}

func NewWorkShop(name string, instructor instructor.Instructor) Workshop {
	w := Workshop{}
	w.Name = name
	w.Instructor = instructor
	return w
}

// we are making the method SignUp as per the interface. (check for interfaces/interfaces.go)
func (w Workshop) SignUp() bool {
	return true
}

package course

import (
	"fmt"

	"mod9/instructor"
)

type Duration float32 //in hours

type Course struct {
	Id         int
	Name       string
	Slug       string
	Legacy     bool
	Duration   Duration
	Instructor instructor.Instructor
}

// this is Stringer method. It will be called implicitly by the 'fmt'. it's like kind of a string representation of your structure, that's kind of idea. So, we can set the output of the struct we have in the string format. works as toString() method in java and other languages.
func (c Course) String() string {
	return fmt.Sprintf("----%v------\n by ------%v------", c.Name, c.Instructor.FirstName)
}

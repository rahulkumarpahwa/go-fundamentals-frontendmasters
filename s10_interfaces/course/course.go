package course

import (
	"mod10/instructor"
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

// we are making the method SignUp as per the interface. (check for interfaces/interfaces.go)
func (c Course) SignUp() bool {
	return true
}

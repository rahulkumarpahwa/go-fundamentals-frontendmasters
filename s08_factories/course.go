package data


type Duration float32 //in hours

type Course struct {
	Id       int
	Name     string
	Slug     string
	Legacy   bool
	Duration Duration
	Instructor Instructor
}

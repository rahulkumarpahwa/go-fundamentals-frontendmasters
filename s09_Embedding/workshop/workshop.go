package workshop
import (
	"mod9/course"
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
	Date time.Time
}

// this embedding will make the properties / datatype available in the Workshop from the course with taking the property of the course.
// see working in the main.go
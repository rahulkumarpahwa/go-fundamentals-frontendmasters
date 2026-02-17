package data

import "fmt"

// here we will create the method related to struct type.
// we can create that in the same file of the struct or in this file different but they need to have the same package.

// note: the name of the method firstletter should be capital.
func (i Instructor) PrintInstructor() string {

	return fmt.Sprintf("Name : %v %v\nId : %v", i.FirstName, i.LastName, i.Id)

}

package data

import "fmt"

var Countries [10]string

var Slice []int

var Codes map[int]string

// init is used to initialise things when app starts.
// we can have the multiple inits functions. they are the only one, which can be repeated.
// they are executed before the main function.
func init() {
	Countries[0] = "argentina"
	Countries[1] = "Spain"
	Countries[4] = "USA"
}

func init() {
	Countries[9] = "India"
	fmt.Println(len(Countries))
	Codes = make(map[int]string)
	Codes[34] = "Apple"
}

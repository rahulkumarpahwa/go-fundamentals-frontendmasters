package main

import "fmt"

var thisispackagevariable string = "this is package variable we can access this in same package any where."

/*
this method prints the data on the terminal and takes only data of the string type.
*/
func printData(data string) {
	fmt.Println("The Data is : ", data)
}

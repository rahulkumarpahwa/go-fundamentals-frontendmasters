package main

import "fmt"
import "mod1/data"

// global variable aka package variable
var url string = "https://rahulkumarpahwa.me"

func main() {

	// local variable
	var message string = "this is the message from the go course"

	msg := "Hello world"

	print("Hello from module.")
	print(`hello \n`) // raw literal strings, means what ever type will be shown directly ie. \n will be shown without giving the new line etc.
	print(" hello \n")
	print(message, msg)

	fmt.Println("This is the print statement")

	// we don't need to import when we the two files have the same package.
	// When the compiler is compiling the app, it will first go over all the go files andat the end, all the files that are part of the same package,it's like they are merged into one big file, into one large file.
	//So a package is the only thing that the compiler sees the code sees packages,not files.We separate content in files just for our organization, not forthe compiler, not for the language.
	// The only restriction is they must of course contain the same declaration onthe top, and they should be in the same folder.
	printData("Hello") // like this one, no need to be imported.

	// similarly, we can access the package variable anywhere we want.
	fmt.Println(thisispackagevariable)

	// using the imported variable from another package:
	fmt.Println("Variable from the data package: ", data.Apple)

	// calling the numbers() method 
	numbers()

	// calling the collection.go in data package
	

}

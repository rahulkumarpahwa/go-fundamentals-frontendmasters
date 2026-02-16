package main

import (
	"fmt"
)

func controls() {
	user := make([]int, 0) // slice of length  0
	if user != nil {
		fmt.Println(user)
	} else {
		fmt.Println("user is nil")
	}

	// we can also create the varaible in the if-statement as well. this message variable in only accesiable in the if-else block and not outside.
	if message := "hello msg from control"; user != nil {
		fmt.Println(message)
	} else {
		// do something
	}

	// switch case:
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday!")
		fallthrough // it is opposite of break keyword.
	case "Tuesday":
		fmt.Println("It's Tuesday")
	default:
		fmt.Print("Enter the correct day!")
	}

	// usage of switch with any bool conditions:
	// switch { // no need to provide the condition here.
	// case user == nil:
	// 	fmt.Println("User is nil")
	// case user.active == false:
	// 	fmt.Println("User is not active")
	// }

	// for loop:

	// Classic for loop -
	// for i:=0; i<len(collection); i++ {
	// }

	// For range, similar to "for in" in JS
	// for index := range collection {
	// }

	// For range, similar to "foreach"
	// for key, value := range map {
	// }

	// for loop as the while loop:
	endOfGame := false

	for endOfGame {
		// process Game Loop
	}

	// similarly
	count := 1
	for count < 10 {
		count += 1
	}

	for { // infinite loop

	}
}

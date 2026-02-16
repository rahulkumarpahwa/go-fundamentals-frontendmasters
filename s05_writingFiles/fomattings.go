package main

import "fmt"

func formatting() {
	value := 34.4447878
	// now to convert that in string, we can't use the string(value), as this will give error. instead we will use the fmt.Sprintf() and we can also get the float format to the 2 decimal values.
	stringValue := fmt.Sprintf("%.2f", value)
	fmt.Println(stringValue)
}

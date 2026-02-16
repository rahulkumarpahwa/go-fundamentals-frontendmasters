package main

import "fmt"

func save() {

}

func save2(text string) {
	fmt.Println("saving", text)
}

func add(a int, b int) int {

	return a + b
}

// if we have to return the multiple values, then return at once.
func addAndSubtract(a int, b int) (int, int) {

	return a + b, a - b
}

func calculateTax(price float32) (float32, float32, string) {
	return price * 0.2, price * 0.3, "hello"
}

// Panic :
func printIntValue(num int) {
	if num > 10 {
		panic("this is the panic.") // The panic built-in function stops normal execution of the current goroutine.
	}
}

func main() {

	// defer : this will be always called at the end. multiple defer will create the stack and follows LIFO
	defer fmt.Println("this is line printed using the defer keyword. this will be printed at last after the last one defer, as this forms the stack.")
	defer fmt.Println("this is line printed using the defer keyword. this will be printed at last.")

	// having multiple values rrturn from the function and handling them.
	cityTax, stateTax, _ := calculateTax(100)
	fmt.Println(cityTax, stateTax)

	// using the birthday method to get the concept of pointers
	age := 34
	fmt.Println(age)
	nextBirthday(age) // now the age should increase by 1.
	fmt.Println(age)  // still 34, because the value is pass by value and not reference.

	// this works as the value id pass by reference.
	age2 := 34
	fmt.Println(age2)
	nextBirthdayPointer(&age2) // now the age should increase by 1.
	fmt.Println(age2)          // still 34, because the value is pass by reference and not value.

	// here %v specifier can be used for any type, normally %d is for integer.
	fmt.Printf("The age without the reference usage is %v and with reference is %v\n", age, age2) // it does not add new line.

	// we also have the concept of the double pointers or the pointers to the pointers.
	age3 := 34
	age3pointer := &age3
	nextBirthdayPointerPointer(&age3pointer)
	fmt.Println(age3)

	//panic:
	printIntValue(9)

	// control structures:
	controls()

}

package main

import "fmt"

func main() {
	var operation string
	var num1, num2 int

	fmt.Println("Calculator GO 1.0")
	fmt.Println("==================")

	fmt.Println("Which Operation do you want? (add, subtract, multiply, divide)")
	fmt.Scan(&operation)
	// fmt.Scanf("%s", &operation)
	fmt.Println("Enter the first Number?")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter the second Number")
	fmt.Scan(&num2)

	switch operation {
	case "+":
		fmt.Println("The result will be : ", num1+num2)
	case "-":
		fmt.Println("Result : ", num1-num2)
	case "/":
		fmt.Println("Result : ", num1/num2)
	case "*":
		fmt.Println("Result : ", num1*num2)
	default:
		fmt.Println("Enter the valid operation")
	}

}

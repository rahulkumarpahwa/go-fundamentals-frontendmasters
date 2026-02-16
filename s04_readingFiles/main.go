package main

import (
	"fmt"
	"os"
)

func readTextFiles(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("We can't read the file.")
		// return "" // when error happen, we want to handle that, then we return the "" empty string.
		// panic("AWWAAAAASWWWWWWWWWWWaaaa") //other way we can panic with a message.
		return "", err // handling the error using the design pattern ie. returning one data and one error.
	} else {
		// here we are converting the bytes array (content) to string using the conversion by UTF8 method.
		return string(content), nil
	}

}

func main() {
	rootPath, _ := os.Getwd()
	fmt.Println(rootPath)

	content, err := readTextFiles("./text.txt")
	if err != nil {
		fmt.Printf("error happened : %v", err)
		panic(err)
	} else {
		fmt.Println(content)
	}

}

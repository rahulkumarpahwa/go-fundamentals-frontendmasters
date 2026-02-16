package main

import ( // this is teh import block. either we can import sinle-single or create this one. similarly we can create the package level varaible block.
	"fmt"
	"os"
)

func writingFile(filename string, content string) error {

	// write the command in terminal : 'go doc os WriteFile' to get the details about the function name.
	err := os.WriteFile(filename, []byte(content), 0644) // last param is the file mode. we use the value 0644 as per the example in go docs.
	// also, here we are using the []byte() to get the byte array from the string.
	return err
}

func main() {
	readFile, _ := os.ReadFile("./text.txt")
	newString := fmt.Sprintf("original content : %v \nnew content, double of the original : %v\n%v ", string(readFile), string(readFile), string(readFile)) // this method returns the formated string like the Printf.
	writingFile("./newFile.txt", newString)
	// this will create the new file in the same folder name newFile.txt and write the content to it.

	// formatting:
	formatting()
}

package main

import "fmt"

func nextBirthday(birthday int) {
	birthday = birthday + 1
}

func nextBirthdayPointer(birthday *int) {
	fmt.Println(birthday)
	*birthday = *birthday + 1
}

// we can have the pointer to the pointer as well. 
func nextBirthdayPointerPointer(birthday **int) {
	**birthday = **birthday + 1
}

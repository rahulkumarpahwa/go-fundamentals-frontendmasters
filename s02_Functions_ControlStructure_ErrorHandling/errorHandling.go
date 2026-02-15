package main

// this is the design pattern of handling the value and the error at the same time.
func readUser(id int) (user any, err any) {
	// we proceed with reading and see a bool ok value
	ok := true
	if ok {
		return user, nil
	} else {
		return nil, err
	}
}

// func main(){
// id := 89
// user, err :=readUser(id)
// }

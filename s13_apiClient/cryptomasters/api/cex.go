package api

import "mod13/datatypes" // importing the package

// here we will have the api related functions and variables.
// we have imported the package and will get the things from it which are exported only.

func GetRate(currency string) (datatypes.Rate, error) {
// go has the inbuilt http in the standard library to get the data from the internet and so we don't need to import any package.
// http package works with both http and https request.


}

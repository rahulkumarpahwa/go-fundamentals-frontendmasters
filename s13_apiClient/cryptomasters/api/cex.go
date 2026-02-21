package api

import (
	"fmt"
	"mod13/datatypes" // importing the package
	"net/http"
	"strings"
)

// here we will have the api related functions and variables.
// we have imported the package and will get the things from it which are exported only.

// const apiURL = "https://cex.io/api/ticker/BTC/INR" // make the apiURL small case so to make it private to the package.
const apiURL = "https://cex.io/api/ticker/%s/%s" // make the apiURL small case so to make it private to the package.
// now, we need to pass the two - 3 letter ISO code of the currency to get the current price of the crypto in that currency. eg. INR for indian rupee and BTC for bitcoin
// to make this dynamic we will put the %s in the apiURL where we have the currency to put on the runtime.

func GetRate(currency string) (datatypes.Rate, error) {
	// go has the inbuilt http in the standard library to get the data from the internet and so we don't need to import any package.
	// http package works with both http and https request.
	currencyUpper := strings.ToUpper(currency) // giving the new string not modifying the original one string.

	res, err := http.Get(fmt.Sprintf(apiURL, currencyUpper))
	// it is a function not a method. it looks like the method of some struct / type. but it is the function of the http package.

	// here, we will use Sprintf() fxn of the fmt package to get the dynamic values in the string in place of the format specifiers.

	// here we need the currency to contain the string in UPPERCASE, so we will need a method over here but currency string does not have any method to convert upper Case. so, we will either create the struct of our own and then create the method over the top of that which will convert to uppercase or in this case, we can use the global method like len(). so we have the strings package, which has function

}

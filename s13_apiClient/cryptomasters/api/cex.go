package api

import (
	// "errors"
	"fmt"
	"io"
	"mod13/datatypes" // importing the package
	"net/http"
	"strings"
)

// here we will have the api related functions and variables.
// we have imported the package and will get the things from it which are exported only.

// const apiURL = "https://cex.io/api/ticker/BTC/INR" // make the apiURL small case so to make it private to the package.
const apiURL =  "https://open.er-api.com/v6/latest/%s"
// I am using the api URL because the api in the lecture is deprecated.
// This api URL provides the exchange rate to all the currencies which are available on the platform from the currency we pass. 

// const apiURL = "https://cex.io/api/ticker/%s" // make the apiURL small case so to make it private to the package.
// now, we need to pass the two - 3 letter ISO code of the currency to get the current price of the crypto in that currency. eg. INR for indian rupee and BTC for bitcoin
// to make this dynamic we will put the %s in the apiURL where we have the currency to put on the runtime.

func GetRate(currency string) (*datatypes.Rate, error) {
	// go has the inbuilt http in the standard library to get the data from the internet and so we don't need to import any package.
	// http package works with both http and https request.
	currencyUpper := strings.ToUpper(currency) // giving the new string not modifying the original one string.

	res, err := http.Get(fmt.Sprintf(apiURL, currencyUpper))
	// it is a function not a method. it looks like the method of some struct / type. but it is the function of the http package.

	// here, we will use Sprintf() fxn of the fmt package to get the dynamic values in the string in place of the format specifiers.

	// here we need the currency to contain the string in UPPERCASE, so we will need a method over here but currency string does not have any method to convert upper Case. so, we will either create the struct of our own and then create the method over the top of that which will convert to uppercase or in this case, we can use the global method like len(). so we have the strings package, which has function 'ToUpper()'

	if err != nil {
		// return datatypes.Rate{}, err // we can return the 'nil' value in case of returning Rate so we will return empty Rate.
		// so, to return the 'nil' we will return the pointer to the Rate and then we can return nil.
		return nil, err // netwwork error, domain error, server down error, URL wrong, etc.
	}
	if res.StatusCode == http.StatusOK {
		// here the data comes, it does not come in one go. It come in chunks. The TCP stack is giving you the response in chunks as soon as they are available from the network.
		// so, one way is to handle it line by line manually till the end of the file when we get the signal that it does not have more or the other way is use the 'io' which has the method ReadAll. ReadAll will receive a reader with a stream and it will read everything until the end, synchronously in one execution.
		// this means thread will held there until the end.
		//But we know that if you wanna get over that we have Go routines, andwe will get there later.But I mean that the thread will wait there. So this is like in JavaScript having an await.
		bodyBytes, err := io.ReadAll(res.Body) //So this is like in JavaScript having an await.
		// It will be there waiting for the response andread all returns on a slice of bytes or error.
		// this err is of this block.
		if err != nil {
			return nil, err // wifi, or offline between the transfer error.
		}
		// As a JSON string just by calling the string converter.It receives the bytes and it will convert that into a string using UTF eight.
		json := string(bodyBytes)
		fmt.Println(json)
	} else { // when server response other than 2xx.
		// here we need to create the error ourself in case of the response other than 200. so, we will use the "errors" package. this has the fxn 'New' which will create the error with msg we pass.
		// but we can also use the fmt.Errorf to create the formatted error.
		return nil, fmt.Errorf("Status Code recieved : %v", res.StatusCode)
	}

	rate := datatypes.Rate{Currency: currencyUpper, Price: 20} // for now we are taking the temporaray value of 20 for the price.
	return &rate, nil                                     // we have to return the pointer of type Rate so use the &.
}

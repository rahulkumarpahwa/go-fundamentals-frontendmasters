package main

import (
	"fmt"
	"mod15/api"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	currencies := []string{"inr", "usd", "eur", "gbp"}
	for _, currency := range currencies {
		wg.Add(1)
		// go GetCurrencyData(currency)
		// wg.Done() // one way is to done here, but this will not wait for the GetCurrencyData to get completed and get executed immediately and nothing gets printed out.

		go func() { // we move to here from the GetCurrencyData. So now we create a new goroutine with this function that will synchronously (only inside) execute,get currency data, and when you finish it will call Done.
			GetCurrencyData(currency)
			wg.Done()
		}() // immediately excuted.

		// so, either pass the waitgroup wg in the method (see last lecture 14 code) or other way is to create the anonymous function (What's a lambda function?A function that we create on the fly without a name and then execute it immediately like the IIFE in js, here we call it as an immediate execution literal function.)  which we will make the go routine by putting the go in front and then pass the GetCurrencyData() and wg.Done() in it. the wg works inside as the go supports the closure concept.
		// now this will make sure that the wait is done properly so as to finish the goroutine created.
		// also, a side note that, now it's not waiting like three seconds,four seconds, the waiting group is working and we get the output immediately.

		// q: Can we refactor it to a new function since we are sending via params?
		// We can, yeah, in that case we won't have the wg variable. So you need to find a way to pass that wg as an argument ordefine that as a package variable, as a global variable.Because here we are taking advantage of we can call wg done becausewe are also using a closure.
	}

	wg.Wait()
}

func GetCurrencyData(currency string) {
	rate, err := api.GetPrice(currency)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(rate.Currency, rate.Price, rate.List["USD"])

}

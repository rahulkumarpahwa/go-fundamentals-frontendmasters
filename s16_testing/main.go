package main

import (
	"fmt"
	"mod16/api"
	"sync"
)

// testing is default in go.
// any file with suffix _test.go is a test file.
// put in the same package.

func main() {
	var wg sync.WaitGroup
	currencies := []string{"inr", "usd", "eur", "gbp"}
	for _, currency := range currencies {
		wg.Add(1)
		go func() {
			GetCurrencyData(currency)
			wg.Done()
		}()
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

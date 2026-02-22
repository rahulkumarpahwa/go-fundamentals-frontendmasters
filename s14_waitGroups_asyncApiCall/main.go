package main

import (
	"fmt"
	"mod14/api"
	"sync"
	// "time"
)

// new Go Routine.
func GetCurrencyData(currency string, wg *sync.WaitGroup) {
	fmt.Println("O") // printed after T
	data, err := api.GetPrice(currency)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data.Price, data.Currency, data.List["USD"])
	wg.Done()
}

// main GoRoutine
func main() {
	var wg sync.WaitGroup
	currencies := []string{"inr", "usd", "eur", "gbp"}
	for _, currency := range currencies {
		wg.Add(1)
		go GetCurrencyData(currency, &wg)
		fmt.Println("T") // printed first
	}
	wg.Wait()

	// so, we will see how we can do non blocking api call which is called the async call.

	// GetCurrencyData("inr")
	// GetCurrencyData("gbp")
	// now, these will get the currency price one by one. synchronous.
	// now, to make then async, then just add the prefix 'go' and then each call will run in its goroutine. ie. its own thread, so they are not blocking the main thread.

	// now if we run directly then nothing will print as the main shutdown before any other goroutine will get completed.
	// go GetCurrencyData("inr")
	// go GetCurrencyData("gbp")
	// go GetCurrencyData("eur")

	// so, to handle it we need add some blocking to the main. this blocking,
	// easiest way is to use the Sleep.
	// time.Sleep(time.Second * 2)
	// note : that the time depend upon the what type of api call is made and how much time does that take.
	// also, the order of the result is not line wise buy the depend upon which goroutine gets completed first. it is different everytime.
	// we can use the channels, here but the channels get into deadlock and to not have deadlock we need to close them but we don't know when to close.
	// so, to solve these problems we will create the 'sync group'.
	// So we are going to create a variable.And typically we're going to create, we call that wg, waiting group.
	// So it's kind of a value that will wait for Goroutines to finish.
	// see line 19.
	// now we will create the slice which will contain the name of the currencies whose price we want to get and will create the for loop over the method GetCurrencyData() for each of the value and now we will pass the 'wg' waitgroup in the method as well.

	//So the idea is that the waitGroup, It's a value. It's kind of a counter actually. So we're going to count, it's a counter,we are adding to that counter, and we are subtracting to the counter.And at any time we can ask that waiting group to wait ie. wg.Wait(); and Wait for the counter to get back to 0, okay? And it will synchronously wait until it gets to 0. So it is like asleep, but not with a fixed time.

	// It will be fixed to that particular moment, okay, when the counter goes to 0. So what's the point?I should talk to my working group and I will add 1 to thatcounter every time we are opening a new Goroutine.
	// and then we will pass the wg as reference (pointer) so that changes made in the original value not in the copy and we will call the wg.Done(); over the waitGroup so as to decrease the value of the wg counter.
	// So this is counting 1, and done is subtracting 1, and then the wait we wait for getting to 0.
	// Do you understand the idea?So every time we open a Goroutine, we say 1, add 1,and every time that Goroutine ends, we should call done.

	// Definitely not O for sure. So what's going on here is when we have a Go call, we are creating a new Goroutine. So basically, the O will be executed in a new Goroutine.And immediately, the main Goroutine will go to line 26 andit will print immediately T, which means, forget about the T and the O.


	// now, question is why are we passing the 'wg' in the GetCurrencyData()?
	// this is because if we call the wg.Done() in the main in the for loop after the fxn call, it will immediately call the Done() method and will not wait for the gorotuine we call to get completed.
	// so, that why we have to pass the wg in the GetCurrencyData and then make it Done() there so as to, make the waitgroup wg wait till the goroutine we created gets completed.
}

// So, back in the cex.go, when you're doing the http.get,I guess I'm used to structures where you have to then do,like, a .then or do an await kind of thing. So is that already built in, that it's waiting until it gets something back?
// -> So you're used to async operations using solutions such as promises, or callbacks.
// Well, you don't have that here. So here this is a synchronous operation, not an asynchronous operation.
// -> Yes. This is a synchronous operation, not an asynchronous operation. When you call get, there is a for waiting for the network to finish, and then it's giving you back the execution. So it is sync, not async.

// Do we have asynchronous operations in Go?
// -> Typically no, you use Goroutines for doing that to avoid blocking the thread.

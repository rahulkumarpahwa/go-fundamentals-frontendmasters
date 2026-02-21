package main

import (
	"fmt"
	"time"
)

// 1.  we have the two goroutine amd we have to transfer data between them then how can we do that. one way is to use the local/package variable which will help them transfer data between these, which is not a good Idea.

// so, we have the channels, which is kind of special variable which is wrapper over the another value.

//so, a main /other  goroutine can create a channel and also wait for the channel. Wait means the goroutine will sleep until the value is recieved.

// also, channel can be buffered and non buffered.

func printMessage(value string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(value)
		time.Sleep(800 * time.Millisecond) // task is taking its time. we are using this to show that.
	}
	channel <- "Message passed through channel from this method / goroutine to main goroutine."
}

func goroutineMethod(msg string, logs chan string) {
	//do some work
	time.Sleep(100 * time.Millisecond)
	logs <- msg
	logs <- msg + "again!"
}

func main() {
	var channel chan string
	channel = make(chan string)

	go printMessage("Frontend Masters", channel)

	msgFromChanneltoMainChannel := <-channel // this is blocking and make the main function wait until the value is recieved.
	fmt.Println(msgFromChanneltoMainChannel)

	// we can use the channel either to send data from one goroutine to the other goroutine or use it to make the one or main goroutine wait until the other goroutine so some event.

	//Note :  Channel is just a variable, so the inter-package usage has the same rule for the normal variables.

	//2. Buffered Channels :
	// a channel where we define how many values we want to send over the channel.
	logs := make(chan string, 2) // here 2 is the size of the buffer channel.
	fmt.Println(logs)
	go goroutineMethod("Hello from main!", logs)
	fmt.Println(<-logs)
	fmt.Println(<-logs)

	// this buffered will help us to send the 2 values before the channel gets wait for the recieving the 2 values.

	// problems : when the main goroutine stops, all the goroutine other than main gets stopped.
	// and we can get into deadlocks.
	// so, to avoid the deadlock, we have to close the channels before ending the programm with 'close'.
	close(channel) // closing the channel 'channel'.
	close(logs) // closing the channel 'logs'.
	// but the problem is when to close the channel and for that we will see to have an example of http server in the next project.

}

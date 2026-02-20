package main

import (
	"fmt"
	"time"
)

// suppose we have the two goroutine amd we have to transfer data between them then how can we do that. one way is to use the local/package variable which will help them transfer data between these, which is not a good Idea.

// so, we have the channels, which is kind of special variable which is wrapper over the another value.

//so, a main /other  goroutine can create a channel and also wait for the channel. Wait means the goroutine will sleep until the value is recieved.

// also, channel can be buffered and non buffered.

func printMessage(value string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(value)
		time.Sleep(800 * time.Millisecond) // task is taking its time. we are using this to show that.
	}
	channel <- "Message passed through channel from this method"
}

func main() {
	var channel chan string
	channel = make(chan string)

	go printMessage("Frontend Masters", channel)

	msgFromChanneltoMainChannel := <-channel // this is blocking and make the main function wait until the value is recieved.
	fmt.Println(msgFromChanneltoMainChannel)

}

package main

import (
	"fmt"
	"mod13/api"
)

// we will call the crypto apis and get the price of the crypto currencies.
// here we will see how to consume data from the internet and later how to send data.

func main() {
	// now, we will call the function GetRate() and we will get the data out of it.
	//

	getRate, err := api.GetRate("inr")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(getRate)
	fmt.Println(getRate.Price, getRate.Currency)
	// getRate is the pointer to the structure so we can access it using the dot operator like the objects in js and get the properties out of it.
	// next step is to parse the json, as we get the string.
}


// note : 1. Lulu is a firewall that I have to on my Macthat I'm going to disable right now you will see why.Because it's detecting a new app in my Mac trying to get to the network,and let's say it's okay, this is new.We won't allow that.I can say Allow, okay?That means that something's happening, it's going into the network, okay, andsomething will happen also.There is a timeout.So at one point, if the firewall is, and there we have it.We have something that we are receiving.Probably, I don't know how to read that, okay?That's probably my print that I have here.So it went to the network and we received something.If I try again, If Lulu is againgiving me an alert, I say Allow, okay?Now, I try again.LuLu gave me an alert.Why?Lulu, my firewall is giving me, I'm saying allow, and that should stay forever.And did I change the rule?Why is Lulu giving me the same alert on every time I run in the app?Does anyone know why?Go is a binary, and actually we are running.But if you remember, running is not actually like running the same app.It's building a new app every time in a temporary directory that I'm not seeingand running it.Every time I'm executing run, it's building a new app.So from an OS point of view,it's a new app that is trying to get to the network, okay.

// 2. And also now, we got the JSON.Why the JSON wasn't there the first time,it was because Lulu was blocking it for a while and then it goes to a timeout.But we see the JSON as a string, so it's working.
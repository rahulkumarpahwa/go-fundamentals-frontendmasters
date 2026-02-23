package main

import (
	"fmt"
	"net/http"
)

// So just as a reminder, if you're a web developer, with Go,you can target web assembly. So in this case, everything that we have seen so far can be actually exposed to JavaScript. So then you compile the module, and then JavaScript can actually execute the function from your package.

// So we are going to serve a website, we are going to see first how toserve it statically, then how we can do some kind of server-side render.And mix data with HTML, see how we can do that,and
// then we will create a very simple web service with Vanilla Go.Because there are also frameworks andlibraries that you can sit on top of Go to help you creating a more complex solution,but we will stay with the standard libraries.

// I mean, the console will always be there, butwe will be focusing into delivering files and content to a web client, a browser.

// And the HTTP package has a couple of functions like exported functions.And the one I'm going to use is called ListenAndserve, ListenAndserve.

func main() {

	http.HandleFunc("/hey", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the /hey route."))
	})
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from '/' route"))
	})

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
}

// It receives first an address, so the address that you want,this is going to open a port in your operating system forincoming connections over TCP.Typically, you use your IP address or your domain colon andthe port, or you can just use colon and a port, any port will work,any number, a valid port number, by the way.

// In this case, if you use that, you don't need to specify your host oryour IP address because it's going to open that port on any networkconnection that you have in your computer.And the second argument,it's called a handler that we're going to start by saying nil.

// ListenAndServe will actually return a possible error option,error value, that's the right word.So if for some reason, I don't know, for example, the port is already open, orin use, or it's not the port that you can actually open because there are some portsreserved for the OS.

// In that case, you will go here and we can just print line,or fmt.Println The server couldn't,or we couldn't open the server, orError while opening the server.

// So now if I run this, that looks pretty empty, but anyway, if I run this,go run., nothing happens, or it seems that nothing is happening.
// So it's not returning control to the console.So what's going on here?
// So the server is there, it's listening in port 3333. Do we have something?No, are we serving something?No, but the server is there.And actually, the thread is there waiting until the server ends.
//When is the server ending? Never. So actually, if you wanna close it,you need to manually go to the console and press Ctrl+C.

// You need to interrupt the server.By default, servers are there up and running always.That's because that's a server, right?So actually what we need to do is do something else with that,but just with that, we do have a server.So you can see, it's actually pretty simple.
// So what else we can do?Actually, we can start talking to the http.And I'm going to refactor this a little bit.

// There are a couple of functions that start with Handle. So we can handle with a type or with a function.

// Let's start with HandleFunc. Parameters : We define a pattern, what's the pattern?A pattern is actually, for now, let's say the URL.But this is a routing system.It's using a multiplexer that is making routes automatically.We have that without a library.So in this case, we can say that for the URL/,or for the URL /hello, doesn't matter,I'm going to pass a function.
// Second Parameter : In Go, you can pass functions as arguments.We haven't done that so far.And also, you can create literal functions.We did that in the previous example.So I can create a handle if I want to.Do you see the signature there?So it's always the same.It's not extremely complicated.But anyway, you have first something known as a responsewriter that we typically call it w.It's an http ResponseWriter.

// We also have the request.That is a pointer to http.Request.That's the signature from an http handler, so I'm not defining that.

// when we go to localhost:3333/, then it will be handled by this function. HandleFunc's second paramter function.

// we have that request and this is kind of writer to the response.If the order seems reversed, right, it's not a natural ordertypically within think request first and response after, but anyway.

// NOW, to write on the screen we need to use the w.Write() method but it takes the []byte ie. slice of byte, so we need to convert our string into []byte by directly passing it in the []byte("here") as we can do that.

// So this default multiplexer that we have, that server, it's smart enough that if there is no handle fora URL "/hello" that I'm trying, it will return a 404, page not found.

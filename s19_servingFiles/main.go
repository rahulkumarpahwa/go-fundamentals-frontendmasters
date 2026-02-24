package main

import (
	"fmt"
	"mod19/auth"
	"net/http"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/auth", auth.HandleAuth) // passing fxn as param

	// File Server:
	fs := http.FileServer(http.Dir("./public"))
	// this will serve all the files in public folder. pass here the relative path.

	server.Handle("/", fs) // see below the chatgpt explanation.

	err := http.ListenAndServe(":1999", server)
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}

}

// what we need to do is to access a utilityfunction from the HTTP package that will let us create a file server.

// A file server will automatically serve to the user all the files in the file system. like the live server extension in vs code.

// Be careful the folder you pick for that.We don't want user to get access to our whole file system.

// So for that, we create that object,it's http.FileServer, and we need to specify the root.

//And for that, we're going to just talk to our http package andwe'll ask for a directory.And the directory will be the public folder that we have here.

//So that's a kind of a handle that will automatically look into the filesystem for every request to see if the file is there or not. So we don't need to handle this manually like reading the URL, go into the filesystem, see if it's there or not, because it's a pretty common pattern.

// You do that directly from the framework.
// In Node.js, you don't have this automatic thing, for example. You have it in Express.js, in a static thing, but that's the framework. eg. .ejs files.

// So then we are going to add another handle. (HandleFunc)
// And we're going to say that, for the root folder, I will use fs as my handle.
// But in this case, it's (fs variable) not a function, it's a type.
// So instead of using HandleFunc, we're going to use just Handle.
// So there are two ways to handle routes.

// Here, with sub nukes that we have in Go, we can parse Handle as a function or we can parse a Type that should confirm to an interface, which means that type (struct) has the method of the interface and so it follows the interface.

// That means that that Type (fs variable) should have a method that at the end is a function.

// ChatGPT Explanation:
/*
1️⃣ Two Ways to Handle Routes in Go

In Go’s net/http, there are two ways to register a route:

HandleFunc → pass a function

Handle → pass a type that implements an interface

Your code already shows the first one.

✅ 1. Using HandleFunc (Function Based)
server.HandleFunc("/auth", auth.HandleAuth)

Here:

/auth is the route

auth.HandleAuth is a function

That function must match this signature:

func(w http.ResponseWriter, r *http.Request)

So here you are directly passing a function as a handler.

This is simple and very common.

✅ 2. Using Handle (Type Based)

Now look at this part:

fs := http.FileServer(http.Dir("./public"))

Here’s what’s important:

http.FileServer() does not return a function

It returns a type

That type implements the http.Handler interface

What is http.Handler?

It is an interface:

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

So any type that has:

ServeHTTP(w http.ResponseWriter, r *http.Request)

automatically satisfies the http.Handler interface.

🔎 What is fs then?
fs := http.FileServer(http.Dir("./public"))

fs is of type:

http.Handler

It is NOT a function.

It is a type that internally has:

ServeHTTP()

method implemented.

🔁 So Why Use Handle Instead of HandleFunc?

Because:

HandleFunc() expects a function

Handle() expects a type implementing http.Handler

So when using FileServer, we must do:

server.Handle("/", fs)

(not shown in your snippet, but that’s what should happen)

Because fs is a type, not a function.
*/

// So, okay, and in this case, we're using that version where we will have the fs as type which has already the method to handle thr route so we will use the Handle instead of HandleFunc to pass the fs in it, otherwise in the normal cases where we create the method we will pass the function to handle that route.

// now, we can open the chrome and go to localhost:1999/ then our static files will be served.  (note : make sure to have the route "/" only in the Handle() method. any new path will lead to addition of that in finding the files from that folder as well relative the Dir passed in the Dir() method)

//Is this server-side rendering?No, this is statically serving a folder.

//By the way, can we call more than one time handler, HandlerFunc? 
// Yes, as many as you want.And actually, the system will look to the pattern (route path pattern) and we'll pick the first one that matches like most of the routing systems.

// Now, if you look at the page source in the HTML in public folder, you only have one article (an H2, a paragraph and an image) in the HTML.
// The rest of the articles are actually coming from JavaScript.
// So it's a client-side render solution.

// JavaScript is downloading a JSON, andit's parsing the JSON and is injecting HTML on the fly.

// But that's not so performant, because we need the user has to wait for JavaScript to load to then go and grab a JSON And then when we have the JSON, it needs to generate that HTML dynamically.

// Important : 
// But now, we are setting that this is bad for performance, sowe can do server-side rendering.
// That is, maybe we can actually inject in the HTML all the articles directly from an aesthetic collection that you have somewhere.

// It can be a database, a JSON, whatever.Makes sense?So I want to create all these articles server-side,so the browser will receive just HTML.
// We can even get rid of JavaScript at all.
// So we can get rid of the script.js and see the same thing on the screen.
// next lecture.

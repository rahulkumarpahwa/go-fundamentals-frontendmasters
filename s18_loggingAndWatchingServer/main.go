package main

import (
	"fmt"
	"mod18/auth"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	// server2 := http.NewServeMux() // similar can be created.

	server.HandleFunc("/auth", auth.HandleAuth)

	// // HTTP default handler : HandleFunc
	// http.HandleFunc("/auth", auth.HandleAuth)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello from go server at '/' route!!!"))
	// })

	// err := http.ListenAndServe(":3333", nil)
	err := http.ListenAndServe(":3333", server)
	if err != nil {
		fmt.Printf("Error :%v", err)
	}
}
 // 1. introduction again:

// if we update the text in w.Write(), we will not the get the update on the localhost:port when refresh.
// this is because, we need to restart the server because the server was compiled and in memories being executed. so, we need to stop the server and run again.
// Okay, so you can see it's pretty simple how you create a server.I mean you can create little functions like this but most of the time,we create functions outside or even in a different package.
// So I'm going to move that function outside,when I move the function to the outside in I need to set a name for the function,so let's call that HandleHello, and then I need to remove the parenthesis now andI will pass, HandleHello as the second argument.
// Be very careful here. HandleAuth, has no codes.I'm not executing the function.I'm passing the function as an argument.That's possible in Go. So if you have a function, you can pass a function as an argument.
// In this case, you are passing the function to the HTTP default handler.

// 2. Creating Server with GoRoutine:

// In this case, you can see we are handling directly over the HTTP package.

// It's kind of a global server, right?So it's me, it's weird the first time you see that I know creating a server I'mjust like, defining handlers for the whole package, which forms a lot of pressure,is okay,

// but technically, because you can open go routines.

// You could have more than one web server running at the same time in different goroutines, not in the same one because it's complicated when you're listen forone, you cannot execute more code.

// So after this line you are there waiting.

// So I cannot open another one in the next line because it will never be executed.

// But I know I can use go routines.
// Well, in that case right now,I'm actually adding handlers to some kind of a global scope. I can change that (http.ListenAndServe(":3333", nil)), and for that is the second argument that I didn't even explain what that was.
// It was called the handler.
// So instead of working globally,I can actually create something that makes things simpler.
// So I can create an object, so I can call this a server.
// So I can create an object, so I can call this a server And the object, it's known as a ServMux.
// It's a multiplexer.
// So for that, we use a factory.
// Do you remember what a factory is?It's just a function starting with the name new.
// So it's a new. We have a NewServMux, like that.
// (see line no.10)

// now instead of creating the HandleFunc with the http ie. http.HandleFunc(), we can create this with server ie. server.HandleFunc() and we can pass the server as the second argument in http.ListenAndServe(":3333", server)

// It's working in the same way.We haven't changed anything more than now. It's better coding. So we are adding the handlers to one special server, andthen we can have another one with other handlers.and run again, it's still working.

// By the way, this thing of stopping and running seems like a lot of work, okay.
// Is Go giving me any other solution, actually, no.

// We don't have any other solution from go, okay?
// So you need to stop because why?Because go is building.
// So if you make a change, we need to make a new bill.

// So all the tools that are available are not from go, are from the community or even from other solutions, you can use things like nodemon that is from node.js.

//So if you search nodemon for Go,you will find here you have five waves to live reload your Go applications, okay.But this is not from Go.
// So these tools are actually detecting your file system.And when there is a new.go file or something you have changed, it's stopping.
// The server and running the server again, okay?But it's a manual thing, but with a tool.
//There is one, for example, it's pretty old.
//It's one of the first one that is called gin.

//It's this one library, low utility for a Go, okay.

// It's a simple command-line utility for live reloading Go web applications.
//You install it and then instead of using run, you use gin run, and gin is like a listener instead of server that is observing the file system and every timeyou change it, it will stop the execution, and build it and run it again.


// 3. Next Step:

// Typically, the next step is to serve a folder, like for example,I could have a public folder here.

//Isn't that that every folder is a package?No, actually no.If you don't have Go files inside, it's not a package.
// How do you know, you don't know it's your organization.
// So if you have folders, without go files, it's not the package.
// In fact, the compiler will just ignore that completely.
// So it's not like every folder will must be a go bucket, NOPE.

// here we can actually add files. Yeah, and I can create an index HTML, images and we can try to serve that.
// I already have, we are not going to waste time writing HTML code.
// So I already have a website ready to serve that is actually here.

// It's a website.
// If you're doing a back end, you don't care about the front end.
// So you just do I know, someone is giving me a public folder,I just need to serve that folder, that's our own responsibility.
// ------------------------------



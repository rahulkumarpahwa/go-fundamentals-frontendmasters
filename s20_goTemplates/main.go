package main

import (
	"fmt"
	"html/template"
	"mod20/auth"
	"net/http"
)

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./template/template.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // header must be send then then write something to the server. otherwise we will get : Write Header error
		w.Write([]byte("Internal Server Error"))
		return
	}

	html.Execute(w, "Hello from the go file in the template") // we can pass here only one value either a single string, a struct or a slice or any other type.
}

func main() {
	server := http.NewServeMux()

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	server.HandleFunc("/auth", auth.HandleAuth)

	server.HandleFunc("/template", handleTemplate)

	err := http.ListenAndServe(":1999", server)
	if err != nil {
		fmt.Println("error : ", err.Error())
	}

}

//1.  Go templates :
// we are going to use something known as Go Templates.
// Go includes a template framework, a template is a format that weuse our strings with some kind of a little language inside,
// if you're using Angular, for example, you know what a template is,when you define the template, your HTML file is actually a template.
// You can insert a special tags that are not the HTML within that template,if you're using React, JSX is kind of a template. It(JSX)'s not really formally a template but kind of, and in Go,it's not specifically targeting web pages, you can use it for any string.

// This is similar, in case you have ever played with PHP, vanilla PHP without the framework, without the templating system.

//You write the dot PHP file andyou mix with HTML with PHP code in the middle.

// 2. Working:
// Anyway, so in this case,we're going to create an HTML file with some Go code in the middle.

// Some, it's a template actually, so it's limited, it's restricted to some ideas.

// It's actually in the HTML package, andthe template can include expressions using double curly braces.

// So if you are coming from Angular,you will feel at home because that's how interpolation works in Angular.

// So you will see this in action in a second.

// Also, there are some options where you can trim spaces,I will show you what that means because it's not completely clear.

// We can mix actions and pipelines as well, this is more complex, we will just cover the basics, the fundamentals, but you can call functions and you can use pipes to define how the output will look like.

// On Angular you also have pipes on templates that are kind of similar.

// examples in angular :
// So for example, you wanna render a price, well, you can parse that through A pipethat will add the currency, the character, like the dollar sign or the euro sign.And then you can parse for another pipe that can format the number with twodigits, and then you can go for another pipe that can convert into Bitcoin.

// we can have the pipeline which will parse that data, the input for different processes.
// Also, we have if-else and range for loops.
// And you can also call functions that you have on your Go file.

// 3. Making:
// Now, we will create the template. for that we will create folder 'template' and put the template inside it. the extension is nothing (or may be gotmpl or tmpl) and just a file. we will then go to the index.html file and copy it in the template file (just name it template).

// Also, we will set the langauge type of the file in the vscode using the bottom bar / footer to the Go templating file.

// remove the script tag and we have the <article/> tag here which we will use as a template.

// so, Idea is to open the double {{}} and inside this we can put the expression.
// Typically, what you're going to do is you're going to receive an object that you use here as your,let's say it's kind of your source of information, so.

/*
import "html/template"
type Person struct {
   Name string
}
func main() {
   t := template.New("my-template").Parse(`
       <!DOCTYPE html>
       <title>My Website<</title>
       <h1>{{.Name}}<</h1>
`)
person := Person{ Name: "Jane Doe" }
output, err := t.ExecuteToString(person)
*/

//so, basically we will have the html/template package and in this we have the New Method with which we will create the new template of our name, then when we want to pass the properties in it like the Name above then we will use the .Name and in the double curly braces.
// At last, we when we execute to string means convert to string then we will pass the struct whose property we want to access and then put in the template.
// Also, while Parsing the html in the backtick we can use the file name instead of writing here.

//In output,we can receive a string in return And you can also write the template directly to the output of a writer or for server. So there are several methods available.

// To work with it, we will create another HandleFunc for the path/ route as : "/template" and we will also create the handlerMethod which is handleTemplate and pass in it.

// later we'll replace that directly with the index HTML.
// So you will remove, we will delete the index HTML from the Public folder andwe will just serve the template version.

// Now, in the handleTemplate() method we will first create the new template with template.ParseFiles() and pass the path here of the template.tmpl file.
// then we will check of err, if no nil and then we will write the internal server error to the header and write in text as well.
//Note:
// Write Header error :
//By the way, if you look at the console, you will have a Write Header error,it's saying me that, I'm saying the 500 is not working. Why?Because I'm sending the Header after writing content.And on HTTP, it is a buffer, the code goes first.So if you have already sent somebody to the client,you cannot change the HTTP code.Does it make sense?So if you wanna send that status code,you need to do it first before sending any content to the browser.That's how the HTTP protocol works.

// Else, we will Execute the parsed html instead of ExecuteToString(), directtly to the w http.ResponseWriter() with the help of method Execute and then second param is the struct whose properties we to access in the template with {{}} and .
// if, we pass the {{}} in the template file we will get the err on the screen.

// So what if we wanna render that string within the template? It's kind of weird, it dot. Dot is actually the value that you have received as input.You can see well dot is just that.
// what if I wanna parse something else like the second value?
// You cannot, it's only one value, that's the idea.

// if you want to access the multiple values, then you can create a structure with everything you need to parse, oryou can parse in a slice.
// Make sense?Typically, you create an object anda structure that you parse to the template and you render for everything.
// and Then Put the . and then value name you want to parse and in the Excute method pass the struct / slice whose property/value you want to access.

// now, go to localhost:1999/template and you will see the page with "String" we passed in the Execute method and when checked the source code in the browser we will get that string in it.

// So it's actually in the source code of the HTML, we have that Test stringthat is coming from a Go value, a Go variable, okay, makes sense?
// So that's the basics of a template.
// It's not a big deal, that dot looks weird, but at the end you get used to.

package main

import (
	"fmt"
	"mod20/auth"
	"net/http"
)

func main() {
	server := http.NewServeMux()

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	server.HandleFunc("/auth", auth.HandleAuth)

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




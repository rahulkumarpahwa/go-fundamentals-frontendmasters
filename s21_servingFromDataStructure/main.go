package main

import (
	"fmt"
	"html/template"
	"mod21/data"
	"net/http"
)

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./template/template.tmpl")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	html.Execute(w, data.GetAll()[0])
}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/template", handleTemplate)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":1999", server)
	if err != nil {
		fmt.Println(err)
	}
}

// >> Is dot kind of type?So no, dot is a reference to the object,to the value that you're sending as the second argument of your execute.

// So it's not the type, it's actually the value.In this case, it's a string, so it looks weird.Most of the time you won't be passing a value, a number or a string,you will be passing an object or a slice, so then the dot makes more sense.

// So the next step is to create.We're going to create the data structure forserving all these exhibitions (articles) from this museum (template page).

// way 1:
// if you go in the public folder, in gallery we have data.json,that's the JSON. We could read the JSON and use our marshal solutions to have it back, that's one option.We can do that because we have already been playing with JSON.
// (we will do self here only again at new route "/json".)

// way 2:
// We can just create like static for now, structure in Go with that information.

// Now, let us create the data folder and create the exhibitions.go in it.

// here, if you copy and paste the JSON andtry to convert it into a normal Go based structure?It works, it's not actually a big deal, we can do that.Remember, we can read the JSON it's not a big deal.Or just because this is kind of boring we can goto GitHub (https://github.com/firtman/go-fundamentals/blob/main/femmuseum/data/exhibitions.go), and I already have that for you.

// It's just the JSON but, With, Let's say, Go format.
// And what's exhibition?
// Exhibition is the type that we need to create, I don't have that type yet.

// It looks like JSON, the only problemis that we don't use here quotes.That if you paste the JSON, the JSON has quotes for the key names or the property names, so you need to go and remove them one by one.

// Fun fact :
// And also some weird thing that I'm not sure you will like isthat the trailing comma is mandatory.So if you try to remove the trailing commas,you can see that image has a trailing comma even if there is no other property.And also there is a trailing comma at the end of the collection (struct), it's mandatory.

// question :
// So you always need to have a trailing comma,does anyone think on a reason of why they did that this way?
//  At least the reason we did it at my last job was sothat if someone added other things afterward it wouldn'tcount that line as a line with a change in GitHub.>
// > Yeah, that's one reason.Also because just to avoid the problem of forgetting the comma andgiving you an error, so simplicity.Remember the three goals of Go, make things simple, to execute fast, andcombine fast, this is simplicity.

// In the exhibitions.go file :
// And this is 'list' with lowercase l,I did this on purpose, that means it's not exported.
// Just to show you that we can implement the pattern of, for example,creating a getter.
// So I can create a function GetAll that returns a slice of exhibition values and I'm going to return the list.

// []Exhibition{ , Remember this is the constructor of a slice of exhibition.The constructor angle is always curly braces, soit's not actually a square brackets as an array in JSON, a reminder.We have the type first and then we open curly braces, andthat's the constructor of types in Go.

// now, here in the main.go we will pass the data in the html.Execute(w, <here>) and we will pass only the first struct of the slice using the index 0.

// also, we need to import the data package and then call the function of that which is data.GetAll() over that which will return the slice of the struct.
// but we will pass only the first one with index 0.

// So now the template is receiving as data source one of these exhibition options, values.

// we can access this using the {{ .Title }} where Title being one of the propety of each struct of Exhibition type.

// we will put these in tags where we want the property to be displayed.

// If this works, let's see, I need to stop my server,run the server again and try it.To refresh template, I'm going to see Aristotle with the right image.And that's actually embedded in the HTML, it's not JavaScript.We've removed the JavaScript, it's actually embedded in the HTML. (make sure to have the file server and also remove the js file from template file.)

// This is server-side render.

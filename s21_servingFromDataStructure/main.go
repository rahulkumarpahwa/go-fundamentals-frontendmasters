package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name    string
	Age     int
	Job     string
	Company string
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("/template/template.tmpl")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	rahul := User{Name: "Rahul Kumar", Age: 30, Job: "SDE", Company: "Q"}

	html.Execute(w, rahul)
}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/template", handleTemplate)

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

// way 2:
// We can just create like static for now, structure in Go with that information.


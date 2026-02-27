package main

import (
	"fmt"
	"mod23/api"
	"mod23/data"
	"net/http"
	"text/template"
)

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./template/index.template")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	err = html.Execute(w, data.GetAll())
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}

func main() {
	server := http.NewServeMux()

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	server.HandleFunc("/template", handleTemplate)

	server.HandleFunc("/api/exhibitions", api.Get)

	err := http.ListenAndServe(":3334", server)
	if err != nil {
		fmt.Println(err)
	}
}

// now, we will see how serve this exhibition as an api. but, we will use the vanilla go. no library, but see some in end.

// we will create the two RESTful operations : GET and POST and rest can created self.

// now create the folder 'api' with package name api.

// all the get operation in get.go and post in post.go in api folder and api package.

// we can connect to choice of our database as well using the library.
